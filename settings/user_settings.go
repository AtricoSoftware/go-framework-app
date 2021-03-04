// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
// SECTION-START: Framework
package settings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const userSettingsSettingName = "UserSettings"

// Lazy value
var userSettingsSettingLazy = NewLazyUserSettingSliceValue(func() []UserSetting { return ParseUserSettingsSetting(viper.Get(userSettingsSettingName)) })

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	return userSettingsSettingLazy.GetValue()
}

// SECTION-END

type UserSetting struct {
	Name            string
	Id              string
	Description     string
	Type            string
	Cmdline         string
	CmdlineShortcut string
	DefaultVal      string
	AppliesTo       []string
}

func ParseUserSettingsSetting(setting interface{}) []UserSetting {
	if setting == nil {
		return []UserSetting{
			UserSetting{
				Name:            "Example", // Name of setting - used in const
				Id:              "example", // viper ID, dotted values form groups
				Description:     "Add your own settings here",
				Type:            "string",
				Cmdline:         "example",
				CmdlineShortcut: "e",
				DefaultVal:      "hello",
				AppliesTo:       []string{"root"},
			}}
	}
	results := make([]UserSetting, len(setting.([]interface{})))
	for i, item := range setting.([]interface{}) {
		mapstructure.Decode(item, &(results[i]))
		results[i].appliesToRootOnly()
	}

	return results
}

func (u *UserSetting) appliesToRootOnly() {
	// If contains root, remove others
	if len(u.AppliesTo) > 1 && sliceContains(u.AppliesTo, "root") {
		u.AppliesTo = []string{"root"}
	}
}

func sliceContains(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func (u UserSetting) NameCode() string {
	return strcase.ToCamel(u.Name)
}

func (u UserSetting) LowerName() string {
	return strcase.ToLowerCamel(u.Name)
}
func (u UserSetting) Filename() string {
	return strcase.ToSnake(u.Name)
}

func (u UserSetting) TypeGetter() string {
	switch u.Type {
	case "string":
		return "viper.GetString"
	case "[]string":
		return "viperEx.GetStringSlice"
	case "bool":
		return "viper.GetBool"
	case "int":
		return "viper.GetInt"
	}
	return ""
}

func (u UserSetting) TypeNameAsCode() string {
	return getTypeNameAsCode(u.Type)
}

func getTypeNameAsCode(str string) string {
	if strings.HasPrefix(str, "[]") {
		return getTypeNameAsCode(str[2:]) + "Slice"
	}
	return strcase.ToCamel(str)
}

func (u UserSetting) TypeFlagAdder() string {
	switch u.Type {
	case "string":
		return "viperEx.AddStringSetting"
	case "[]string":
		return "viperEx.AddStringArraySetting"
	case "bool":
		return "viperEx.AddBoolSetting"
	case "int":
		return "viperEx.AddIntSetting"
	}

	panic(errors.New(fmt.Sprintf("type '%s' is not supported as a setting type", u.Type)))
}

func (u UserSetting) AppliesToCmd(cmd string) bool {
	// No commandline, do not apply to any command
	if u.Cmdline == "" {
		return false
	}
	for _, c := range u.AppliesTo {
		if c == cmd {
			return true
		}
	}
	return false
}

// For template
func (u UserSetting) HasPrefix(text string, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}

// Applies to as comma separated list
func (u UserSetting) AppliesToCSL() string {
	return strings.Replace(strings.Join(u.AppliesTo, ", "), "root", "all", 1)
}
