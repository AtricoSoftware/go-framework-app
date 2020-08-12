package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/common"
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/pkg"
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/requirements"
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
		// Get the requirements
		requirements.GetRequirements(settings.TargetDirectory())
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
		"ApplicationName": settings.ApplicationName(),
		"RepositoryPath":  settings.RepositoryPath(),
		"Commands":        settings.Commands(),
		"UserSettings":    settings.UserSettings(),
	}
}

func generateFile(path string, filename string, contents *template.Template, values interface{}) error {
	fullPath := filepath.Join(path, filename)
	fmt.Println("Writing: ", fullPath)
	os.MkdirAll(filepath.Dir(fullPath), 0755)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if comment := getComment(filepath.Base(filename)); comment != "" {
		writer.WriteString(fmt.Sprintf("%s Generated %s by %s %s\n", comment, time.Now().Format("2006-02-01"), pkg.Name, pkg.Version))
	}
	// DEBUG contents.Execute(os.Stdout, values)

	if err = contents.Execute(writer, values); err != nil {
		return err
	}
	return writer.Flush()
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
