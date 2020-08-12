
package files

import "text/template"

// All the simple files
var Files = make([]*template.Template, 0)

func init() {
Files = append(Files, template.Must(template.New(`.gitignore`).Parse(`# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
`)))
Files = append(Files, template.Must(template.New(`azure-pipelines.yml`).Parse(`trigger:
  branches:
    include:
      - refs/tags/*
pr: none

resources:
  containers:
    - container: go-build-agent
      image: mclaren-cloudplatform-docker-dev-local.jfrog.io/golang:1.13
      endpoint: artifactory-cloud-mcp-docker

variables:
  - group: condition-insight-dev

pool:
  name: 'ConditionInsight-Dev'

stages:
  - stage: BUILD
    displayName: Build
    jobs:
      - job: build_and_publish
        container: go-build-agent
        displayName: Build
        steps:
          - script: |
              VERSION="$(git describe --tags --dirty)"
              echo "##vso[task.setvariable variable=version]$VERSION"
              echo Version = $VERSION
            displayName: Set version
          - script: |
              ./build.sh $(version)
            displayName: Build {{.ApplicationName}}
          - task: ArtifactoryGenericUpload@1
            inputs:
              artifactoryService: 'Artifactory Cloud'
              specSource: 'taskConfiguration'
              fileSpec: |
                {
                  "files": [
                    {
                      "pattern": "./release/*.zip",
                      "target": "conditioninsight-dev-local/{{.ApplicationName}}/$(version)/"
                    }
                  ]
                }
              failNoOp: true
            displayName: Publish {{.ApplicationName}} to Artifactory
`)))
Files = append(Files, template.Must(template.New(`build.sh`).Parse(`MODULE="{{.RepositoryPath}}"
export OUTPUT_NAME="{{.ApplicationName}}"
TARGET_DIR=release
TARGET_PLATFORMS="darwin windows linux"

if [[ ! -z "$1" ]]
then
  VERSION=$1
else
  VERSION=$(git describe --tags --dirty)
fi

export CGO_ENABLED=0
export GOARCH="amd64"

# setup details
# built
BUILT_ON=$(date)
BUILT_BY=$(whoami)
# git
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_COMMIT=$(git rev-parse HEAD)

DETAILS="{\"Built\":{\"On\":\"$BUILT_ON\", \"By\":\"$BUILT_BY\"},\"Git\":{ \"Repository\":\"MODULE\",\"Branch\":\"$GIT_BRANCH\",\"Commit\":\"$GIT_COMMIT\"} }"
# Setup ldflags
LDFLAGS="-s -w"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.Version=$VERSION'"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.BuildDetails=$DETAILS'"


mkdir -p $TARGET_DIR
for GOOS in $TARGET_PLATFORMS; do
    export GOOS
    export EXT=""
    if [[ ${GOOS} == "windows" ]]
    then
      export EXT=".exe"
    fi
    export TARGET="$TARGET_DIR/$VERSION-$GOOS-$GOARCH"
    mkdir -p $TARGET
    go build -v -ldflags="$LDFLAGS" -o $TARGET/$OUTPUT_NAME$EXT

done

cd $TARGET_DIR
find . ! -path . -type d |  cut -d "/" -f2 | awk -v name="$OUTPUT_NAME" '{ print name "_" $1 ".zip -r ./" $1 "/"  }' | xargs -L1 zip -j
#find . ! -path . -type d | xargs -L1 rm -rf

`)))
Files = append(Files, template.Must(template.New(`go.mod`).Parse(`module {{.RepositoryPath}}

go 1.14

require (
)
`)))
Files = append(Files, template.Must(template.New(`main.go`).Parse(`package main

import (
	"{{.RepositoryPath}}/cmd"
)

func main() {
	cmd.Execute()
}
`)))
Files = append(Files, template.Must(template.New(`cmd\root.go`).Parse(`package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"{{.RepositoryPath}}/pkg"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var rootCmd = &cobra.Command{
	Use:   pkg.Name,
	Short: pkg.Summary,
	Long:  fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "alternate config file")
}

func initConfig() {
	// Config file
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err := tryReadConfig(); err != nil {
			// Fail if specified config cannot be read
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Standard name for config
		viper.SetConfigName(fmt.Sprintf(".%s", pkg.Name))
		// Try current working directory
		dir, err := os.Getwd()
		if err == nil {
			viper.AddConfigPath(dir)
			err = tryReadConfig()
		}
		if err != nil {
			// Finally, try home directory
			dir, err = homedir.Dir()
			if err == nil {
				viper.AddConfigPath(dir)
				tryReadConfig()
			}
		}
	}
}

func tryReadConfig() error {
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	return err
}
`)))
Files = append(Files, template.Must(template.New(`cmd\todo.go`).Parse(`package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/settings"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "summary",
	Run: func(*cobra.Command, []string) {
		// Implementation here!
		settings := settings.GetSettings() // Get the default settings
		fmt.Printf("Backup = %s\n", settings.Example())
	},
}

func init() {
	settings.AddExampleFlag(exampleCmd.PersistentFlags())
	rootCmd.AddCommand(exampleCmd)
}
`)))
Files = append(Files, template.Must(template.New(`cmd\version.go`).Parse(`package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/pkg"
)

var fullVersion bool

var showVersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Shows version",
	Run: func(*cobra.Command, []string) {
		if fullVersion {
			fmt.Println(pkg.Name)
			fmt.Println(pkg.Description)
		}
		fmt.Println(pkg.Version)
		if fullVersion {
			fmt.Println()
			var details map[string]interface{}
			if err := json.Unmarshal([]byte(pkg.BuildDetails), &details); err == nil && len(details) > 0 {
				fmt.Println("Details")
				fmt.Println("-------")
				displaySection(details, "")
			}
		}
	},
}

func init() {
	showVersionCommand.PersistentFlags().BoolVarP(&fullVersion, "full", "f", false, "Full program information")
	rootCmd.AddCommand(showVersionCommand)
}

func displaySection(section map[string]interface{}, indent string) {
	for k, v := range section {
		fmt.Printf("%s%s:", indent, k)
		switch v.(type) {
		case map[string]interface{}:
			fmt.Println()
			displaySection(v.(map[string]interface{}), indent+"  ")
		default:
			fmt.Printf(" %s\n", v)
		}
	}
}
`)))
Files = append(Files, template.Must(template.New(`pkg\info.go`).Parse(`package pkg

var Name = "{{.ApplicationName}}"
var Summary = "TODO"
var Description = "TODO"

// Set by build.sh
var Version = "development-version"
var BuildDetails = "{}"
`)))
Files = append(Files, template.Must(template.New(`settings\settings.go`).Parse(`package settings

type Settings interface {
	// TODO - Add your own settings as required
	Example() string
}

// Get the settings for this run
func GetSettings() Settings {
	return theSettings{}
}

// Stub object for settings interface
type theSettings struct{}
`)))
Files = append(Files, template.Must(template.New(`settings\todo.go`).Parse(`package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"{{.RepositoryPath}}/viperEx"
)

// This is the name by which the setting is specified on the commandline
const exampleSettingName = "example-setting"

// Fetch the setting
func (theSettings) Example() string {
	return viper.GetString(exampleSettingName)
}

func AddExampleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, exampleSettingName, "Description of setting")
	// Use P version if you want a shorthand option
	//	viperEx.AddStringSettingP(flagSet, exampleSettingName, "e", "Description of setting")
}
`)))
Files = append(Files, template.Must(template.New(`viperEx\add_setting.go`).Parse(`package viperEx

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var flags = make(map[string]*pflag.Flag)

// Used in testing to clear settings
func Reset() {
	viper.Reset()
	flags = make(map[string]*pflag.Flag)
}

func AddBoolSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddBoolSettingP(flagSet, name, "", description)
}

func AddBoolSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.BoolP(name, shorthand, false, description) })
}

func AddStringSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddStringSettingP(flagSet, name, "", description)
}

func AddStringSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.StringP(name, shorthand, "", description) })
}

func AddStringArraySetting(flagSet *pflag.FlagSet, name string, description string) {
	AddStringArraySettingP(flagSet, name, "", description)
}

func AddStringArraySettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.StringArrayP(name, shorthand, []string{}, description) })
}

func addSetting(flagSet *pflag.FlagSet, name string, createFlag func()) {
	if flag, ok := flags[name]; ok {
		// TODO [Improvement] - check type is the same
		// Add existing flag
		flagSet.AddFlag(flag)
	} else {
		// Create new flag
		createFlag()
		flag = flagSet.Lookup(name)
		// Bind to viper
		if err := viper.BindPFlag(name, flag); err != nil {
			log.Fatal("Unable to bind flag:", err)
		}
		// Store for next time
		flags[name] = flag
	}
}
`)))

}