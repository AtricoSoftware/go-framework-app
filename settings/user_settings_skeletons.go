package settings

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func resolveSkeleton(setting UserSetting) UserSetting {
	// Read in skeleton files (first time only)
	ReadSkeletonFiles(theSettings{}.SkeletonFiles())
	if setting.Skeleton != "" {
		// Unknown skeleton
		if skeleton, ok := SkeletonCloset[strings.ToLower(setting.Skeleton)]; ok {
			// Update skeleton with changes
			for fieldIdx, defaultValue := range UserSettingFieldIndexes {
				// Get setting value
				value := reflect.ValueOf(setting).Field(fieldIdx)
				// Check against default (ie not set)
				if !reflect.DeepEqual(value.Interface(), defaultValue) {
					// Set "changed" value
					reflect.ValueOf(&skeleton).Elem().Field(fieldIdx).Set(reflect.ValueOf(setting).Field(fieldIdx))
				}
			}
			return skeleton
		} else {
			fmt.Fprintln(os.Stderr, "Unknown skeleton: ", setting.Skeleton)
		}
	}
	return setting
}

var UserSettingType reflect.Type
var UserSettingFieldIndexes = make(map[int]interface{})

func init() {
	setting := UserSetting{}
	UserSettingType = reflect.TypeOf(setting)
	for fieldIdx := 0; fieldIdx < UserSettingType.NumField(); fieldIdx++ {
		// Skip Skeleton itself
		if UserSettingType.Field(fieldIdx).Name != "Skeleton" {
			// Get setting value
			value := reflect.ValueOf(setting).Field(fieldIdx)
			if value.CanInterface() {
				defaultValue := reflect.New(value.Type()).Elem().Interface()
				UserSettingFieldIndexes[fieldIdx] = defaultValue
			}
		}
	}
}

func ReadSkeletonFiles(files []string) {
	if !skeletonsRead {
		for _, file := range files {
			conf := viper.New()
			conf.SetConfigFile(file)
			if err := conf.ReadInConfig(); err == nil {
				GetVerboseService().VerbosePrintln("Reading skeleton file: ", file)
				var setting UserSetting
				for name, item := range conf.AllSettings() {
					if err = mapstructure.Decode(item, &setting); err == nil {
						setting.Skeleton = name
						SkeletonCloset[name] = setting
					} else {
						fmt.Fprintf(os.Stderr, "Could not parse skeleton '%s': %v\n", name, err)

					}
				}
			} else {
				fmt.Fprintf(os.Stderr, "Could not open skeleton file '%s': %v\n", file, err)
			}
		}
		skeletonsRead = true
	}
}

var skeletonsRead = false
var SkeletonCloset = map[string]UserSetting{
	"dry-run": {
		Name:        "DryRun",
		Id:          "DryRun",
		Description: "Show actions but do not perform them",
		Type:        "bool",
		Cmdline:     "dry-run",
		EnvVar:      "DRY_RUN",
	},
	"quiet": {
		Name:            "Quiet",
		Id:              "Quiet",
		Description:     "Suppress stdout",
		Type:            "bool",
		Cmdline:         "quiet",
		CmdlineShortcut: "q",
	},
	"recursive": {
		Name:            "Recursive",
		Id:              "Recursive",
		Description:     "Operate recursively",
		Type:            "bool",
		Cmdline:         "recursive",
		CmdlineShortcut: "R",
	},
	"force": {
		Name:            "Force",
		Id:              "Force",
		Description:     "Force operation",
		Type:            "bool",
		Cmdline:         "force",
		CmdlineShortcut: "f",
	},
	"directory": {
		Name:            "Directory",
		Id:              "Directory",
		Description:     "Directory",
		Type:            "string",
		Cmdline:         "directory",
		CmdlineShortcut: "d",
		DefaultVal:      ".",
	},
	"file": {
		Name:            "File",
		Id:              "File",
		Description:     "File",
		Type:            "string",
		Cmdline:         "file",
		CmdlineShortcut: "f",
	},
}
