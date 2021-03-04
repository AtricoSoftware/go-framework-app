package file_writer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type GeneratedFileInfo interface {
	BaseDir() string
	OriginalPath() string
	BackupPath() string
	FullOriginalPath() string
	FullBackupPath() string
}

type generatedFileInfo struct {
	baseDir      string
	backupPath   string
	originalPath string
}

func (i generatedFileInfo) BaseDir() string          { return i.baseDir }
func (i generatedFileInfo) OriginalPath() string     { return i.originalPath }
func (i generatedFileInfo) BackupPath() string       { return i.backupPath }
func (i generatedFileInfo) FullOriginalPath() string { return filepath.Join(i.baseDir, i.originalPath) }
func (i generatedFileInfo) FullBackupPath() string   { return filepath.Join(i.baseDir, i.backupPath) }

func GenerateFile(targetDir string, fileTemplate FileTemplate, values TemplateValues) (GeneratedFileInfo, error) {
	return GenerateNamedFile(targetDir, fileTemplate, "", values)
}

func GenerateNamedFile(targetDir string, fileTemplate FileTemplate, name string, values TemplateValues) (GeneratedFileInfo, error) {
	var info generatedFileInfo
	info.baseDir = targetDir
	info.originalPath = safeSprintf(fileTemplate.Path, name)
	os.MkdirAll(filepath.Dir(info.FullOriginalPath()), 0755)
	fmt.Println("Writing: ", info.FullOriginalPath())
	fileExists := fileExists(info.FullOriginalPath())
	// Backup existing file
	if fileExists {
		// Backup
		info.backupPath = fmt.Sprintf("%s_%s.bak", info.OriginalPath(), values["BackupSuffix"])
		os.Rename(info.FullOriginalPath(), info.FullBackupPath())
	}
	// DEBUG contents.Execute(os.Stdout, values)
	// Write to buffer first
	buffer := new(bytes.Buffer)
	var err error
	if err = fileTemplate.MainFile.Execute(buffer, values); err == nil {
		newfileContents := buffer.Bytes()
		if fileTemplate.FileTemplateType == MixedTemplate && fileExists {
			// Get required imports
			isGoFile := filepath.Ext(info.OriginalPath()) == ".go"
			importList := make([]ImportItem, 0)
			if isGoFile {
				importList = AddImports(newfileContents, importList)
			}
			// Strip buffer into sections
			newFile := make(map[string]string)
			unusedSections := make(map[string]interface{})
			for _, part := range StripSections(newfileContents) {
				if part.Section != "" {
					newFile[part.Section] = part.Contents
					unusedSections[part.Section] = nil
				}
			}
			// Read existing file
			existContents, _ := ioutil.ReadFile(info.FullBackupPath())
			if isGoFile {
				importList = AddImports(existContents, importList)
				unusedSections[ImportsSection] = nil
				newFile[ImportsSection] = FormatImports(importList)
			}
			existingFile := StripSections(existContents)
			// Write file, replacing sections
			var file *os.File
			if file, err = os.Create(info.FullOriginalPath()); err == nil {
				defer file.Close()
				for _, part := range existingFile {
					if cont, ok := newFile[part.Section]; ok {
						file.WriteString(cont)
						delete(unusedSections, part.Section)
					} else {
						file.WriteString(part.Contents)
					}
				}
			}
			for sect, _ := range unusedSections {
				fmt.Fprintf(os.Stderr, "WARNING: Missing section - %s\n", sect)
			}
		} else {
			// Simply copy to file
			err = ioutil.WriteFile(info.FullOriginalPath(), buffer.Bytes(), 0644)
		}
	}
	return info, err
}
