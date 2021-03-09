package file_writer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type generatedFileInfo struct {
	baseDir      string
	backupPath   string
	originalPath string
}

func (i generatedFileInfo) fullOriginalPath() string { return filepath.Join(i.baseDir, i.originalPath) }
func (i generatedFileInfo) fullBackupPath() string   { return filepath.Join(i.baseDir, i.backupPath) }

func (f fileWriter) GenerateFile(fileTemplate FileTemplate, values TemplateValues) error {
	return f.GenerateNamedFile(fileTemplate, "", values)
}

func (f fileWriter) GenerateNamedFile(fileTemplate FileTemplate, name string, values TemplateValues) error {
	var info generatedFileInfo
	info.baseDir = f.config.TargetDirectory()
	info.originalPath = safeSprintf(fileTemplate.Path, name)
	os.MkdirAll(filepath.Dir(info.fullOriginalPath()), 0755)
	fmt.Println("Writing: ", info.fullOriginalPath())
	fileExists := fileExists(info.fullOriginalPath())
	// Backup existing file
	if fileExists {
		// Backup
		info.backupPath = fmt.Sprintf("%s_%s.bak", info.originalPath, values["BackupSuffix"])
		os.Rename(info.fullOriginalPath(), info.fullBackupPath())
	}
	// DEBUG contents.Execute(os.Stdout, values)
	// Write to buffer first
	buffer := new(bytes.Buffer)
	var err error
	if err = fileTemplate.MainFile.Execute(buffer, values); err == nil {
		newfileContents := buffer.Bytes()
		if fileTemplate.FileTemplateType == MixedTemplate && fileExists {
			// Get imports
			isGoFile := filepath.Ext(info.originalPath) == ".go"
			importList := make([]ImportItem, 0)
			if isGoFile {
				importList = AddImports(newfileContents, importList)
			}
			// Get requirements
			isModFile := filepath.Base(info.originalPath) == "go.mod"
			requirementsList := make([]RequireItem, 0)
			if isModFile {
				requirementsList = AddRequirements(newfileContents, requirementsList)
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
			existContents, _ := ioutil.ReadFile(info.fullBackupPath())
			if isGoFile {
				importList = AddImports(existContents, importList)
				unusedSections[ImportsSection] = nil
				newFile[ImportsSection] = FormatImports(importList)
			}
			if isModFile {
				requirementsList = AddRequirements(existContents, requirementsList)
				unusedSections[RequiresSection] = nil
				newFile[RequiresSection] = FormatRequirements(requirementsList)
			}
			existingFile := StripSections(existContents)
			// Write file, replacing sections
			var file *os.File
			if file, err = os.Create(info.fullOriginalPath()); err == nil {
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
			err = ioutil.WriteFile(info.fullOriginalPath(), buffer.Bytes(), 0644)
		}
	}
	if err == nil {
		*f.generatedFiles = append(*f.generatedFiles, info)
	}
	return err
}
