// Generated 2021-02-24 16:58:12 by go-framework v1.5.0
package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/file_writer"
	"github.com/AtricoSoftware/go-framework-app/resources"
	"github.com/AtricoSoftware/go-framework-app/settings"

	"github.com/spf13/viper"
)

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

	generatedFiles := make([]file_writer.GeneratedFileInfo, 0)

	// Create all standard files
	for _, t := range resources.Files {
		if info, err := file_writer.GenerateFile(svc.config.TargetDirectory(), t.Name(), t, values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
	}
	// Create commands/api
	cmdPath := filepath.Join(svc.config.TargetDirectory(), "cmd")
	apiPath := filepath.Join(svc.config.TargetDirectory(), "api")
	for _, command := range svc.config.Commands() {
		values["Command"] = command
		if info, err := file_writer.GenerateFile(cmdPath, fmt.Sprintf("%s.go", command.Name), resources.Templates["cmd"], values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
		// Do not overwrite existing api (this is what the user will change)
		file_writer.GenerateFileIfNotPresent(apiPath, fmt.Sprintf("%s.go", command.Name), resources.Templates["api"], values)
	}
	// Create settings
	settingsPath := filepath.Join(svc.config.TargetDirectory(), "settings")
	lazyTypes := make(map[string]settings.UserSetting, 0)
	for _, setting := range svc.config.UserSettings() {
		values["Setting"] = setting
		if info, err := file_writer.GenerateFile(settingsPath, fmt.Sprintf("%s.go", setting.Filename()), resources.Templates["setting"], values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
		if setting.TypeGetter() == "" {
			// Custom setting (no overwrite)
			file_writer.GenerateFileIfNotPresent(settingsPath, fmt.Sprintf("%s_impl.go", setting.Filename()), resources.Templates["setting_impl"], setting)
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
		if info, err := file_writer.GenerateFile(settingsPath, "lazy_implementations.go", resources.Templates["lazy_implementations"], settings); err == nil {
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
	// Get the requirements
	for _, url := range getRequirements(svc.config.Libraries()) {
		GoCommand(svc.config.TargetDirectory(), "get", url)
	}
	// Clean up the files
	GoCommand("fmt", "./...")
	// Remove backups with no changes
	file_writer.CleanupBackups(generatedFiles)
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

func getRequirements(userLibraries []string) []string {
	libs := make(map[string]string, len(resources.Requirements)+len(userLibraries))
	// System requirements, then user libraries
	for _, entry := range append(resources.Requirements, userLibraries...) {
		url, ver := splitUrlVersion(entry)
		libs[url] = ver
	}
	return mergeUrlVersion(libs)
}

func splitUrlVersion(raw string) (url, version string) {
	parts := strings.Split(raw, "@")
	l := len(parts)
	if l > 0 {
		url = parts[0]
	}
	if l > 1 {
		version = parts[1]
	}
	return url, version
}

func mergeUrlVersion(urlVersions map[string]string) []string {
	urls := make([]string, 0, len(urlVersions))
	for url, ver := range urlVersions {
		var newUrl string
		if ver != "" {
			newUrl = fmt.Sprintf("%s@%s", url, ver)
		} else {
			newUrl = url
		}
		urls = append(urls, newUrl)
	}
	return urls
}
