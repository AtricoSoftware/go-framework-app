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
const templatesPkg = "templates"

func main() {
	// Create files templates
	createTemplates(filesPkg, fileHeader, fileFooter, addFilesTemplate)
	// Create specific templates
	createTemplates(templatesPkg, templatesHeader, templatesFooter, addTemplatesTemplate)
}

func createTemplates(pkg string, header string, footer string, addTemplate func(name string, contents []byte)string) {
	// Create templates from files
	tFile, err := os.Create(filepath.Join(pkg, "templates.go"))
	if err != nil {
		panic(err)
	}
	defer tFile.Close()
	// Write file header
	tFile.WriteString(header)
	// Find all files starting with underscore
	filepath.Walk(pkg, func(path string, info os.FileInfo, err error) error {
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
			newName := strings.Replace(path, info.Name(), info.Name()[1:], 1)[len(pkg)+1:]
			fmt.Println("Adding: ", newName)
			// TODO - handle ` in file (readme)
			tFile.WriteString(addTemplate(newName, contents))
		}
		return nil
	})
	// Write file footer
	tFile.WriteString(footer)
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

func addFilesTemplate(name string, contents []byte) string {
	return fmt.Sprintf("Files = append(Files, template.Must(template.New(`%s`).Parse(`%s`)))\n", name, contents)
}

var templatesHeader = fmt.Sprintf(`
package %s

import "text/template"

// Specific file templates
var Templates = make(map[string]*template.Template)

func init() {
`, templatesPkg)

var templatesFooter = `
}`
func addTemplatesTemplate(name string, contents []byte) string {
	name2 := strings.Replace(filepath.Base(name), filepath.Ext(name), "", 1)
	return fmt.Sprintf("Templates[`%s`] = template.Must(template.New(`%s`).Parse(`%s`))\n", name2, name2, contents)
}
