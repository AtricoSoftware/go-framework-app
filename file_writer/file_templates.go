package file_writer

import (
	"fmt"
	"strings"
	"text/template"
)

type FileTemplate struct {
	FileTemplateType
	Path     string
	MainFile *template.Template
	Sections map[string]*template.Template
}

type FileTemplateType int

const (
	FrameworkTemplate FileTemplateType = iota
	MixedTemplate     FileTemplateType = iota
)

func (f FileTemplateType) String() string {
	switch f {
	case FrameworkTemplate:
		return "FrameworkTemplate"
	case MixedTemplate:
		return "MixedTemplate"
	}
	panic("invalid template type")
}

func ParseTemplateType(str string) (ftype FileTemplateType, ok bool) {
	ftype = FrameworkTemplate
	switch strings.ToLower(str) {
	case "framework":
		ok = true
	case "mixed":
		ftype = MixedTemplate
		ok = true
	}
	return ftype, ok
}

func MustParseTemplateType(str string) FileTemplateType {
	if ftype, ok := ParseTemplateType(str); ok {
		return ftype
	}
	panic(fmt.Sprintf("invalid template type: %s", str))
}
