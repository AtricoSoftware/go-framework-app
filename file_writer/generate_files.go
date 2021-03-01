package file_writer

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/AtricoSoftware/go-framework-app/pkg"
)

type GeneratedFileInfo struct {
	backupPath   string
	originalPath string
	comment      bool
}

func GenerateFile(path string, filename string, contents *template.Template, values interface{}) (GeneratedFileInfo, error) {
	return GenerateFileImpl(path, filename, true, contents, values)
}
func GenerateFileIfNotPresent(path string, filename string, contents *template.Template, values interface{}) (GeneratedFileInfo, error) {
	return GenerateFileImpl(path, filename, false, contents, values)
}
func GenerateFileImpl(path string, filename string, overwrite bool, contents *template.Template, values interface{}) (info GeneratedFileInfo, err error) {
	info.originalPath = filepath.Join(path, filename)
	os.MkdirAll(filepath.Dir(info.originalPath), 0755)
	if fileExists(info.originalPath) {
		if !overwrite {
			fmt.Println("Skipping: ", info.originalPath)
			return info, errors.New("file already exists")
		}
		info.backupPath = fmt.Sprintf("%s_%s.bak", info.originalPath, runTime.Format("2006-01-02_15-04-05"))
		os.Rename(info.originalPath, info.backupPath)
	}
	fmt.Println("Writing: ", info.originalPath)
	var file *os.File
	if file, err = os.Create(info.originalPath); err == nil {
		defer file.Close()
		writer := bufio.NewWriter(file)
		comment := getComment(filepath.Base(filename))
		if comment != "" {
			info.comment = true
			commentStr := fmt.Sprintf("Generated %s by %s %s", runTime.Format("2006-01-02 15:04:05"), pkg.Name, pkg.Version)
			writer.WriteString(fmt.Sprintf(comment, commentStr))
			writer.WriteString("\n")
		}
		// DEBUG contents.Execute(os.Stdout, values)

		if err = contents.Execute(writer, values); err == nil {
			err = writer.Flush()
		}
	}
	return info, err
}
