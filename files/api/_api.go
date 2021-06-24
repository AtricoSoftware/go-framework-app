package api

import (
	"{{.RepositoryPath}}/settings"
)

type Api interface {
{{- range .Commands}}
	// {{.Description}}
	{{.ApiName}}(config settings.Settings) error
{{- end}}
}

// Get the API for normal run
func GetApi() Api {
	return theApi{}
}

// Stub object for api interface
type theApi struct{}