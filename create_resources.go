// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const filesPkg = "files"

func main() {
	// Create file templates
	tFile, err := os.Create(filepath.Join(filesPkg, "templates.go"))
	if err != nil {
		panic(err)
	}
	defer tFile.Close()
	// Write file header
	tFile.WriteString(fileHeader)
	// Find all files starting with underscore
	filepath.Walk(filesPkg, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.HasPrefix(info.Name(), "_") {
			// Read file contents
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			contents, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			// Strip leading underscore and filesPkg
			newName := strings.Replace(path, info.Name(), info.Name()[1:], 1)[len(filesPkg)+1:]
			fmt.Println("Adding: ", newName)
			// TODO - handle ` in file (readme)
			tFile.WriteString(fmt.Sprintf("Files = append(Files, template.Must(template.New(`%s`).Parse(`%s`)))\n", newName, contents))
		}
		return nil
	})
	// Write file footer
	tFile.WriteString(fileFooter)
	tFile.Sync()
}

var fileHeader = fmt.Sprintf(`
package %s

import "text/template"

// All the simple files
var Files = make([]*template.Template, 0)

func init() {
`, filesPkg)

var fileFooter = `
}`

func fileAddTemplate(path string, contents string) string {
	return fmt.Sprintf(`			
Files = append(template.Must(template.New(%s).Parse(contents)\n")

`, path)
}
