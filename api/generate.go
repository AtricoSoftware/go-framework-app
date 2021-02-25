// Generated 2021-02-24 16:58:12 by go-framework development-version
package api

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/resources"
	"github.com/AtricoSoftware/go-framework-app/settings"

	"github.com/spf13/viper"

	"github.com/AtricoSoftware/go-framework-app/pkg"
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
	values := createTemplateValues(svc.config)

	generatedFiles := make([]generatedFileInfo, 0)

	// Create all standard files
	for _, t := range resources.Files {
		if info, err := generateFile(svc.config.TargetDirectory(), t.Name(), t, values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
	}
	// Create commands/api
	cmdPath := filepath.Join(svc.config.TargetDirectory(), "cmd")
	apiPath := filepath.Join(svc.config.TargetDirectory(), "api")
	for _, command := range svc.config.Commands() {
		values["Command"] = command
		if info, err := generateFile(cmdPath, fmt.Sprintf("%s.go", command.Name), resources.Templates["cmd"], values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
		// Do not overwrite existing api (this is what the user will change)
		generateFileIfNotPresent(apiPath, fmt.Sprintf("%s.go", command.Name), resources.Templates["api"], values)
	}
	// Create settings
	settingsPath := filepath.Join(svc.config.TargetDirectory(), "settings")
	for _, setting := range svc.config.UserSettings() {
		values["Setting"] = setting
		if info, err := generateFile(settingsPath, fmt.Sprintf("%s.go", setting.Filename()), resources.Templates["setting"], values); err == nil {
			generatedFiles = append(generatedFiles, info)
		}
		if setting.TypeGetter() == "" {
			// Custom setting (no overwrite)
			generateFileIfNotPresent(settingsPath, fmt.Sprintf("%s_impl.go", setting.Filename()), resources.Templates["setting_impl"], setting)
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
	for _, info := range generatedFiles {
		if info.backupPath != "" {
			if filesEqual(info.originalPath, info.backupPath, info.comment) {
				// Files equal, remove backup
				os.Remove(info.backupPath)
			}
		}
	}
	return nil
}

var runTime = time.Now()

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

func createTemplateValues(settings settings.Settings) map[string]interface{} {
	return map[string]interface{}{
		"SingleReadConfiguration": settings.SingleReadConfiguration(),
		"ApplicationTitle":        settings.ApplicationTitle(),
		"ApplicationName":         settings.ApplicationName(),
		"ApplicationSummary":      settings.ApplicationSummary(),
		"ApplicationDescription":  settings.ApplicationDescription(),
		"RepositoryPath":          settings.RepositoryPath(),
		"Commands":                settings.Commands(),
		"UserSettings":            settings.UserSettings(),
		"Libraries":               settings.Libraries(),
	}
}

type generatedFileInfo struct {
	backupPath   string
	originalPath string
	comment      bool
}

func generateFile(path string, filename string, contents *template.Template, values interface{}) (generatedFileInfo, error) {
	return generateFileImpl(path, filename, true, contents, values)
}
func generateFileIfNotPresent(path string, filename string, contents *template.Template, values interface{}) (generatedFileInfo, error) {
	return generateFileImpl(path, filename, false, contents, values)
}
func generateFileImpl(path string, filename string, overwrite bool, contents *template.Template, values interface{}) (info generatedFileInfo, err error) {
	info.originalPath = filepath.Join(path, filename)
	os.MkdirAll(filepath.Dir(info.originalPath), 0755)
	if fileExists(info.originalPath) {
		if !overwrite {
			fmt.Println("Skipping: ", info.originalPath)
			return info, errors.New("file already exists")
		}
		info.backupPath = fmt.Sprintf("%s_%s.bak", info.originalPath, runTime.Format("2006-01-02_15-04-05"))
		os.Rename(info.originalPath, info.backupPath)
	}
	fmt.Println("Writing: ", info.originalPath)
	var file *os.File
	if file, err = os.Create(info.originalPath); err == nil {
		defer file.Close()
		writer := bufio.NewWriter(file)
		comment := getComment(filepath.Base(filename))
		if comment != "" {
			info.comment = true
			writer.WriteString(fmt.Sprintf("%s Generated %s by %s %s\n", comment, runTime.Format("2006-01-02 15:04:05"), pkg.Name, pkg.Version))
		}
		// DEBUG contents.Execute(os.Stdout, values)

		if err = contents.Execute(writer, values); err == nil {
			err = writer.Flush()
		}
	}
	return info, err
}

func fileExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return err == nil
}

// True if files are equal
// False if error
func filesEqual(path1, path2 string, skipComment bool) bool {
	if file1, err := os.Open(path1); err == nil {
		defer file1.Close()
		if file2, err := os.Open(path2); err == nil {
			defer file2.Close()
			scanner1 := bufio.NewScanner(file1)
			scanner2 := bufio.NewScanner(file2)
			// Read each line
			for scanner1.Scan() && scanner2.Scan() {
				// Skip first line?
				if skipComment {
					skipComment = false
					continue
				}
				// Compare lines
				if scanner1.Text() != scanner2.Text() {
					return false
				}
			}
			return true
		}
	}
	return false
}

func getComment(filename string) string {
	if filename == ".gitignore" {
		return "#"
	}
	if filename == "go.mod" {
		return "//"
	}
	switch filepath.Ext(filename) {
	case ".go":
		return "//"
	case ".sh", ".yaml", ".yml":
		return "#"
	}
	return ""
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
