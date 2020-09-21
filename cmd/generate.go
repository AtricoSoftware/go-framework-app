package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/common"
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/pkg"
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/templates"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/files"
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/settings"
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

		// Create all standard files
		for _, t := range files.Files {
			generateFile(settings.TargetDirectory(), t.Name(), t, values)
		}
		// Create commands/api
		cmdPath := filepath.Join(settings.TargetDirectory(), "cmd")
		apiPath := filepath.Join(settings.TargetDirectory(), "api")
		for _, command := range settings.Commands() {
			values["Command"] = command
			generateFile(cmdPath, fmt.Sprintf("%s.go", command.Name()), templates.Templates["cmd"], values)
			generateFile(apiPath, fmt.Sprintf("%s.go", command.Name()), templates.Templates["api"], values)
		}
		// Create settings
		settingsPath := filepath.Join(settings.TargetDirectory(), "settings")
		for _, setting := range settings.UserSettings() {
			values["Setting"] = setting
			generateFile(settingsPath, fmt.Sprintf("%s.go", setting.Filename()), templates.Templates["setting"], values)
		}
		// Copy generator settings if found (for future reference)
		data, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err == nil {
			configFile := filepath.Base(viper.ConfigFileUsed())
			destination := filepath.Join(settings.TargetDirectory(), configFile)
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
	},
}

func init() {
	settings.AddTargetDirectoryFlag(generateCmd.PersistentFlags())
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
		"ApplicationName":        settings.ApplicationName(),
		"ApplicationSummary":     settings.ApplicationSummary(),
		"ApplicationDescription": settings.ApplicationDescription(),
		"RepositoryPath":         settings.RepositoryPath(),
		"Commands":               settings.Commands(),
		"UserSettings":           settings.UserSettings(),
		"Libraries":              settings.Libraries(),
	}
}

func generateFile(path string, filename string, contents *template.Template, values interface{}) error {
	fullPath := filepath.Join(path, filename)
	fmt.Println("Writing: ", fullPath)
	os.MkdirAll(filepath.Dir(fullPath), 0755)
	backupFile(fullPath)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if comment := getComment(filepath.Base(filename)); comment != "" {
		writer.WriteString(fmt.Sprintf("%s Generated %s by %s %s\n", comment, runTime.Format("2006-01-02 15:04:05"), pkg.Name, pkg.Version))
	}
	// DEBUG contents.Execute(os.Stdout, values)

	if err = contents.Execute(writer, values); err != nil {
		return err
	}
	return writer.Flush()
}

// Backup file (if it already exists)
func backupFile(fullPath string) {
	source, err := os.Open(fullPath)
	if !os.IsNotExist(err) {
		defer source.Close()
		// Copy to backup file
		destination, err := os.Create(fmt.Sprintf("%s_%s.bak", fullPath, runTime.Format("2006-01-02_15-04-05")))
		if err == nil {
			defer destination.Close()
			io.Copy(destination, source)
		}
	}
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
