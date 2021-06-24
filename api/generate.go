// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
// SECTION-START: Framework
package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/atrico-go/container"
	"github.com/spf13/viper"

	"github.com/AtricoSoftware/go-framework-app/api/file_writer"
	"github.com/AtricoSoftware/go-framework-app/resources"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type GenerateApi Runnable
type GenerateApiFactory Factory

type generateApiFactory struct {
	container.Container
}

func (f generateApiFactory) Create(args []string) Runnable {
	RegisterApiGenerate(f.Container)
	var theApi GenerateApi
	f.Container.Make(&theApi)
	return theApi
}

// SECTION-END

func RegisterApiGenerate(c container.Container) {
	file_writer.RegisterFileWriter(c)
	c.Singleton(func(config settings.Settings, verboseService settings.VerboseService, fileWriter file_writer.FileWriter) GenerateApi {
		return generateApi{config, verboseService, fileWriter}
	})
}

type generateApi struct {
	settings.Settings
	settings.VerboseService
	fileWriter file_writer.FileWriter
}

// Generate framework app
func (svc generateApi) Run() error {
	validateMandatorySetting(svc.ApplicationName(), "Application name")
	validateMandatorySetting(svc.RepositoryPath(), "Repository path")
	// Ensure target folder exists
	validateFolder(svc.TargetDirectory())
	// Create values for the template
	values := svc.fileWriter.CreateTemplateValues()

	var err error
	// Create all standard files
	for _, t := range resources.Files {
		svc.fileWriter.GenerateFile(t, values)
	}
	// Create commands/api
	for _, command := range svc.Commands() {
		values["Command"] = command
		svc.fileWriter.GenerateNamedFile(resources.Templates["cmd"], command.FileName(), values)
		if !command.NoImplementation {
			svc.fileWriter.GenerateNamedFile(resources.Templates["api"], command.FileName(), values)
		}
	}
	// Create settings
	cachedTypes := make(map[string]settings.UserSetting, 0)
	for _, setting := range svc.UserSettings() {
		values["Setting"] = setting
		svc.fileWriter.GenerateNamedFile(resources.Templates["setting"], setting.Filename(), values)
		if svc.SingleReadConfiguration() {
			cachedTypes[setting.Type] = setting
		}
	}
	// Cache implementations
	if len(cachedTypes) > 0 {
		settings := make([]settings.UserSetting, 0, len(cachedTypes))
		for _, st := range cachedTypes {
			settings = append(settings, st)
		}
		sort.Slice(settings, func(i, j int) bool { return settings[i].TypeNameAsCode() < settings[j].TypeNameAsCode() })
		values["CachedTypes"] = settings
		svc.fileWriter.GenerateFile(resources.Templates["cache_implementations"], values)
	}
	// Copy generator settings if found (for future reference)
	data, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err == nil {
		configExt := filepath.Ext(viper.ConfigFileUsed())
		destination := filepath.Join(svc.TargetDirectory(), fmt.Sprintf(".go-framework%s", configExt))
		ioutil.WriteFile(destination, data, 0644)
	}
	// Module dependencies
	fmt.Println("Getting dependencies (go get)")
	GoCommand(svc.TargetDirectory(), "get", "-u", "all")
	fmt.Println("Downloading dependencies (go mod download)")
	GoCommand(svc.TargetDirectory(), "mod", "download")
	// Format code
	fmt.Println("Formatting code (gofmt)")
	ExecuteCommand(svc.TargetDirectory(), "gofmt", "-w", ".")
	// Build the app
	fmt.Println("Building code")
	const buildFile = "build-temp"
	GoCommand(svc.TargetDirectory(), "build", "-o", buildFile, ".")
	os.Remove(filepath.Join(svc.TargetDirectory(), buildFile))
	// Run the unit tests
	fmt.Println("Running unit tests")
	GoCommand(svc.TargetDirectory(), "test", "./unit-tests")
	// Clean up the files
	fmt.Println("Cleaning up")
	svc.fileWriter.CleanupFiles()
	// Remove backups with no changes
	svc.fileWriter.RemoveObsoleteBackups()
	return nil
}

func validateMandatorySetting(setting string, name string) {
	if setting == "" {
		fmt.Println(name, " not specified")
		os.Exit(-1)
	}
}

func validateFolder(path string) {
	src, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			panic(err)
		}
	} else {
		if src.Mode().IsRegular() {
			fmt.Println(path, " already exists as a file!")
			os.Exit(-1)
		}
	}
}
