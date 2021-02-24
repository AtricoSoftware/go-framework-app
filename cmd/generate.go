package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/AtricoSoftware/go-framework-app/common"
	"github.com/AtricoSoftware/go-framework-app/pkg"
	"github.com/AtricoSoftware/go-framework-app/templates"

	"github.com/AtricoSoftware/go-framework-app/files"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate framework app",
	Run: func(*cobra.Command, []string) {
		settings := settings.GetSettings() // Get the default settings
		validateMandatorySetting(settings.ApplicationName(), "Application name")
		validateMandatorySetting(settings.RepositoryPath(), "Repository path")
		// Ensure target folder exists
		validateFolder(settings.TargetDirectory())
		// Create values for the template
		values := createTemplateValues(settings)

		generatedFiles := make([]generatedFileInfo, 0)

		// Create all standard files
		for _, t := range files.Files {
			if info, err := generateFile(settings.TargetDirectory(), t.Name(), t, values); err == nil {
				generatedFiles = append(generatedFiles, info)
			}
		}
		// Create commands/api
		cmdPath := filepath.Join(settings.TargetDirectory(), "cmd")
		apiPath := filepath.Join(settings.TargetDirectory(), "api")
		for _, command := range settings.Commands() {
			values["Command"] = command
			if info, err := generateFile(cmdPath, fmt.Sprintf("%s.go", command.Name()), templates.Templates["cmd"], values); err == nil {
				generatedFiles = append(generatedFiles, info)
			}
			// Do not overwrite existing api (this is what the user will change)
			generateFileIfNotPresent(apiPath, fmt.Sprintf("%s.go", command.Name()), templates.Templates["api"], values)
		}
		// Create settings
		settingsPath := filepath.Join(settings.TargetDirectory(), "settings")
		for _, setting := range settings.UserSettings() {
			values["Setting"] = setting
			if info, err := generateFile(settingsPath, fmt.Sprintf("%s.go", setting.Filename()), templates.Templates["setting"], values); err == nil {
				generatedFiles = append(generatedFiles, info)
			}
			if setting.TypeGetter() == "" {
				// Custom setting (no overwrite)
				generateFileIfNotPresent(settingsPath, fmt.Sprintf("%s_impl.go", setting.Filename()), templates.Templates["setting_impl"], setting)
			}
		}
		// Copy generator settings if found (for future reference)
		data, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err == nil {
			configExt := filepath.Ext(viper.ConfigFileUsed())
			destination := filepath.Join(settings.TargetDirectory(), fmt.Sprintf(".go-framework%s", configExt))
			ioutil.WriteFile(destination, data, 0644)
		}
		// Get the requirements
		for url, version := range settings.Libraries() {
			var pkg string
			if version == "" {
				pkg = url
			} else {
				pkg = fmt.Sprintf("%s@%s", url, version)
			}
			common.GoCommand(settings.TargetDirectory(), "get", pkg)
		}
		// Clean up the files
		common.GoCommand("fmt", "./...")
		// Remove backups with no changes
		for _,info := range generatedFiles {
			if info.backupPath != "" {
				if filesEqual(info.originalPath,info.backupPath, info.comment) {
					// Files equal, remove backup
					os.Remove(info.backupPath)
				}
			}
		}
	},
}

func init() {
	settings.AddTargetDirectoryFlag(generateCmd.PersistentFlags())
	settings.AddApplicationTitleFlag(generateCmd.PersistentFlags())
	settings.AddApplicationNameFlag(generateCmd.PersistentFlags())
	settings.AddRepositoryPathFlag(generateCmd.PersistentFlags())
	rootCmd.AddCommand(generateCmd)
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
		"ApplicationTitle":       settings.ApplicationTitle(),
		"ApplicationName":        settings.ApplicationName(),
		"ApplicationSummary":     settings.ApplicationSummary(),
		"ApplicationDescription": settings.ApplicationDescription(),
		"RepositoryPath":         settings.RepositoryPath(),
		"Commands":               settings.Commands(),
		"UserSettings":           settings.UserSettings(),
		"Libraries":              settings.Libraries(),
	}
}

type generatedFileInfo struct {
	backupPath string
	originalPath string
	comment bool
}

func generateFile(path string, filename string, contents *template.Template, values interface{}) (generatedFileInfo, error) {
	return generateFileImpl(path, filename, true, contents, values)
}
func generateFileIfNotPresent(path string, filename string, contents *template.Template, values interface{}) (generatedFileInfo, error) {
	return generateFileImpl(path, filename, false, contents, values)
}
func generateFileImpl(path string, filename string, overwrite bool, contents *template.Template, values interface{}) (info generatedFileInfo, err error) {
	info.originalPath = filepath.Join(path, filename)
	fmt.Println("Writing: ", info.originalPath)
	os.MkdirAll(filepath.Dir(info.originalPath), 0755)
	if fileExists(info.originalPath) {
		if !overwrite {
			return info, errors.New("file already exists")
		}
		info.backupPath = fmt.Sprintf("%s_%s.bak", info.originalPath, runTime.Format("2006-01-02_15-04-05"))
		os.Rename(info.originalPath, info.backupPath)
	}
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
func filesEqual(path1,path2 string, skipComment bool)  bool {
	if file1,err := os.Open(path1); err == nil {
		defer file1.Close()
		if file2,err := os.Open(path2); err == nil {
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
