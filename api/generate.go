// Generated 2021-03-05 09:21:54 by go-framework development-version
// SECTION-START: Framework
package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/atrico-go/container"
	"github.com/spf13/viper"

	"github.com/AtricoSoftware/go-framework-app/file_writer"
	"github.com/AtricoSoftware/go-framework-app/resources"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

// SECTION-END

func RegisterGenerate(c container.Container) {
	c.Singleton(func(config settings.Settings) GenerateApi { return generateApi{config: config} })
}

type generateApi struct {
	config settings.Settings
}

// Generate framework app
func (svc generateApi) Run() error {
	validateMandatorySetting(svc.config.ApplicationName(), "Application name")
	validateMandatorySetting(svc.config.RepositoryPath(), "Repository path")
	// Ensure target folder exists
	validateFolder(svc.config.TargetDirectory())
	// Create values for the template
	values := file_writer.CreateTemplateValues(svc.config)
	// Add comment string/backup suffix to values
	now := time.Now()
	values["Comment"] = file_writer.FileComment(now)
	values["BackupSuffix"] = now.Format("2006-01-02_15-04-05")

	generatedFiles := make([]file_writer.GeneratedFileInfo, 0)

	var info file_writer.GeneratedFileInfo
	var err error
	// Create all standard files
	for _, t := range resources.Files {
		if info, err = file_writer.GenerateFile(svc.config.TargetDirectory(), t, values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
	}
	// Create commands/api
	for _, command := range svc.config.Commands() {
		values["Command"] = command
		for _, pkg := range []string{"cmd", "api"} {
			if info, err = file_writer.GenerateNamedFile(svc.config.TargetDirectory(), resources.Templates[pkg], command.Name, values); err == nil {
				generatedFiles = append(generatedFiles, info)
			}
		}
	}
	// Create settings
	lazyTypes := make(map[string]settings.UserSetting, 0)
	for _, setting := range svc.config.UserSettings() {
		values["Setting"] = setting
		if info, err = file_writer.GenerateNamedFile(svc.config.TargetDirectory(), resources.Templates["setting"], setting.Filename(), values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
		if svc.config.SingleReadConfiguration() {
			lazyTypes[setting.Type] = setting
		}
	}
	// lazy implementations
	if len(lazyTypes) > 0 {
		settings := make([]settings.UserSetting, 0, len(lazyTypes))
		for _, st := range lazyTypes {
			settings = append(settings, st)
		}
		sort.Slice(settings, func(i, j int) bool { return settings[i].TypeNameAsCode() < settings[j].TypeNameAsCode() })
		values["LazySettings"] = settings
		if info, err = file_writer.GenerateFile(svc.config.TargetDirectory(), resources.Templates["lazy_implementations"], values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
	}
	// Copy generator settings if found (for future reference)
	data, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err == nil {
		configExt := filepath.Ext(viper.ConfigFileUsed())
		destination := filepath.Join(svc.config.TargetDirectory(), fmt.Sprintf(".go-framework%s", configExt))
		ioutil.WriteFile(destination, data, 0644)
	}
	// Clean up the files
	file_writer.CleanupFiles(svc.config.RepositoryPath(), generatedFiles)
	// Remove backups with no changes
	file_writer.RemoveObsoleteBackups(generatedFiles)
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
