// Generated 2021-05-24 17:41:23 by go-framework development-version
// SECTION-START: Framework
package api

import (
	"github.com/atrico-go/container"
	"github.com/AtricoSoftware/go-framework-app/settings"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"github.com/AtricoSoftware/go-framework-app/api/file_writer"
	"github.com/AtricoSoftware/go-framework-app/resources"
	"github.com/spf13/viper"
)

type GenerateApi Runnable
type GenerateApiFactory Factory

type generateApiFactory struct {
	container.Container
}

func (f generateApiFactory) Create() Runnable {
	RegisterApiGenerate(f.Container)
	var theApi GenerateApi
	f.Container.Make(&theApi)
	return theApi
}
// SECTION-END

func RegisterApiGenerate(c container.Container) {
	file_writer.RegisterFileWriter(c)
	c.Singleton(func(config settings.Settings, fileWriter file_writer.FileWriter) GenerateApi {
		return generateApi{config, fileWriter}
	})
}

type generateApi struct {
	config     settings.Settings
	fileWriter file_writer.FileWriter
}

// Generate framework app
func (svc generateApi) Run() error {
	validateMandatorySetting(svc.config.ApplicationName(), "Application name")
	validateMandatorySetting(svc.config.RepositoryPath(), "Repository path")
	// Ensure target folder exists
	validateFolder(svc.config.TargetDirectory())
	// Create values for the template
	values := svc.fileWriter.CreateTemplateValues()

	var err error
	// Create all standard files
	for _, t := range resources.Files {
		svc.fileWriter.GenerateFile(t, values)
	}
	// Create commands/api
	for _, command := range svc.config.Commands() {
		values["Command"] = command
		for _, pkg := range []string{"cmd", "api"} {
			svc.fileWriter.GenerateNamedFile(resources.Templates[pkg], command.FileName(), values)
		}
	}
	// Create settings
	cachedTypes := make(map[string]settings.UserSetting, 0)
	for _, setting := range svc.config.UserSettings() {
		values["Setting"] = setting
		svc.fileWriter.GenerateNamedFile(resources.Templates["setting"], setting.Filename(), values)
		if svc.config.SingleReadConfiguration() {
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
		destination := filepath.Join(svc.config.TargetDirectory(), fmt.Sprintf(".go-framework%s", configExt))
		ioutil.WriteFile(destination, data, 0644)
	}
	// Module dependencies
	GoCommand(svc.config.TargetDirectory(), "get", "-u", "all")
	GoCommand(svc.config.TargetDirectory(), "mod", "download")
	// Clean up the files
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
