// +build ignore (Used to create resources only)

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/AtricoSoftware/go-framework-app/api/file_writer"
)

const filesPkg = "files"
const templatesPkg = "templates"

func main() {
	// Create files templates
	createTemplates(filesPkg, fileHeader, fileFooter, addFilesTemplate)
	// Create specific templates
	createTemplates(templatesPkg, templatesHeader, templatesFooter, addTemplatesTemplate)
}

func createTemplates(pkg string, header string, footer string, addTemplate func(templateName string, fileTemplateType file_writer.FileTemplateType, path string, content string) string) {
	// Create templates from files
	fileFolder := filepath.Join("resources", pkg)
	tFile, err := os.Create(filepath.Join("resources", fmt.Sprintf("tmpl_%s.go", pkg)))
	if err != nil {
		panic(err)
	}
	defer tFile.Close()
	// Write file header
	tFile.WriteString(header)
	// Find all files
	filepath.Walk(fileFolder, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.HasPrefix(info.Name(), "_") {
			// Read file contents
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			// Read main content
			header, contents := readToSeparator(file)
			var fileTemplateType file_writer.FileTemplateType
			var ok bool
			if fileTemplateType, ok = file_writer.ParseTemplateType(header["Type"]); !ok {
				fileTemplateType = file_writer.FrameworkTemplate
			}
			if _, err = template.New("").Parse(contents); err == nil {
				// Strip resource path and correct filename (deps on header)
				templateName, newPath := calculateFilename(path[len(fileFolder)+1:], header["Name"])
				tFile.WriteString(addTemplate(templateName, fileTemplateType, newPath, strings.ReplaceAll(contents, "`", "`+\"`\"+`")))
			} else {
				fmt.Fprintf(os.Stderr, "Failed to parse template: %s (%v)\n", info.Name(), err)
			}
		}
		return nil
	})
	// Write file footer
	tFile.WriteString(footer)
	tFile.Sync()
}

type Header map[string]string

// Read upto next separator
// Discard separator
// If EOF reached treat as separator
func readToSeparator(file *os.File) (header Header, content string) {
	header = make(Header)
	headerRead := false
	contentBuilder := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "---" {
			return header, contentBuilder.String()
		}
		if headerRead || json.Unmarshal([]byte(scanner.Text()), &header) != nil {
			contentBuilder.WriteString(scanner.Text())
			contentBuilder.WriteString("\n")
		}
		headerRead = true
	}
	return header, contentBuilder.String()
}

var fileHeader = `package resources

import (
	"text/template"

	"github.com/AtricoSoftware/go-framework-app/api/file_writer"
)

// All the "static" files
var Files = make([]file_writer.FileTemplate, 0)

func init() {
`

var fileFooter = `
}`

func createTemplateInit(fileTemplateType file_writer.FileTemplateType, path string, content string) string {
	return fmt.Sprintf("file_writer.FileTemplate{FileTemplateType: file_writer.%s, Path: `%s`,MainFile: template.Must(template.New(`mainFile`).Parse(`%s`))}", fileTemplateType.String(), path, content)

}

func addFilesTemplate(_ string, fileTemplateType file_writer.FileTemplateType, path string, content string) string {
	return fmt.Sprintf("Files = append(Files, %s)\n", createTemplateInit(fileTemplateType, path, content))
}

var templatesHeader = `package resources

import (
	"text/template"

	"github.com/AtricoSoftware/go-framework-app/api/file_writer"
)

// Specific file templates
var Templates = make(map[string]file_writer.FileTemplate)

func init() {
`

var templatesFooter = `
}`

func addTemplatesTemplate(templateName string, fileTemplateType file_writer.FileTemplateType, path string, content string) string {
	return fmt.Sprintf("Templates[`%s`] = %s\n", templateName, createTemplateInit(fileTemplateType, path, content))
}

func calculateFilename(original string, name string) (templateName, path string) {
	originalBase := filepath.Base(original)
	var base string
	if name == "" {
		// No name specified, simply strip leading underscore
		base = originalBase[1:]
	} else {
		// Name specified...
		if filepath.Ext(name) != "" {
			// Specified name has extension, use it "as is"
			base = name
		} else {
			// Use existing extension
			base = name + filepath.Ext(originalBase)
		}
	}
	templateName = originalBase[1:]
	if ext := filepath.Ext(templateName); ext != "" {
		templateName = strings.Replace(templateName, ext, "", 1)
	}
	return templateName, strings.Replace(original, originalBase, base, 1)
}
