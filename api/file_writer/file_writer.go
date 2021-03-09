package file_writer

import (
	"time"

	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/settings"
)

type FileWriter interface {
	// Create the values for use in templates
	CreateTemplateValues() TemplateValues
	// Generate files
	GenerateFile(fileTemplate FileTemplate, values TemplateValues) error
	GenerateNamedFile(fileTemplate FileTemplate, name string, values TemplateValues) error
	// Format go files
	CleanupFiles()
	// Remove backups of files that have not changed
	RemoveObsoleteBackups()
}

func RegisterFileWriter(c container.Container) {
	c.Singleton(func(config settings.Settings) FileWriter { return fileWriter{config, time.Now(), new([]generatedFileInfo)} })
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type fileWriter struct {
	config         settings.Settings
	now            time.Time
	generatedFiles *[]generatedFileInfo
}
