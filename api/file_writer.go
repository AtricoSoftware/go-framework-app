package api

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	"github.com/AtricoSoftware/go-framework-app/pkg"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type TemplateValues map[string]interface{}

func createTemplateValues(settings settings.Settings) TemplateValues {
	t := reflect.TypeOf(settings)
	numMethods := t.NumMethod()
	values := make(TemplateValues, numMethods)
	for i := 0; i < numMethods; i++ {
		method := t.Method(i)
		values[method.Name] = reflect.ValueOf(settings).MethodByName(method.Name).Call([]reflect.Value{})[0].Interface()
	}
	return values
}

type generatedFileInfo struct {
	backupPath   string
	originalPath string
	comment      bool
}

func generateFile(path string, filename string, contents *template.Template, values interface{}) (generatedFileInfo, error) {
	return generateFileImpl(path, filename, true, contents, values)
}
func generateFileIfNotPresent(path string, filename string, contents *template.Template, values interface{}) (generatedFileInfo, error) {
	return generateFileImpl(path, filename, false, contents, values)
}
func generateFileImpl(path string, filename string, overwrite bool, contents *template.Template, values interface{}) (info generatedFileInfo, err error) {
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

func fileExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return err == nil
}

// True if files are equal
// False if error
func filesEqual(path1, path2 string, skipComment bool) bool {
	if file1, err := os.Open(path1); err == nil {
		defer file1.Close()
		if file2, err := os.Open(path2); err == nil {
			defer file2.Close()
			scanner1 := bufio.NewScanner(file1)
			scanner2 := bufio.NewScanner(file2)
			// Read each line
			for scanner1.Scan() && scanner2.Scan() {
				// Skip first line?
				if skipComment {
					skipComment = false
					continue
				}
				// Compare lines
				if scanner1.Text() != scanner2.Text() {
					return false
				}
			}
			return true
		}
	}
	return false
}

func getComment(filename string) string {
	if filename == ".gitignore" {
		return "# %s"
	}
	if filename == "go.mod" {
		return "// %s"
	}
	switch filepath.Ext(filename) {
	case ".go":
		return "// %s"
	case ".sh", ".yaml", ".yml":
		return "# %s"
	case ".md":
		return "[comment]: <> ( %s )"
	}
	return ""
}
