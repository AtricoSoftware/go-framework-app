package api

import (
	"{{.RepositoryPath}}/settings"
)

// {{.Command.Description}}
func (theApi) {{.Command.ApiName}}(config settings.Settings) error {
	// Implementation here!
	return nil
}