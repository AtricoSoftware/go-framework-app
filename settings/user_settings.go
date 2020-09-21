package settings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

const settingsSettingName = "settings"

type UserSetting interface {
	Name() string
	Description() string
	Type() string
	Cmdline() string
	CmdlineShortcut() string
	DefaultVal() string
	AppliesTo() []string

	Id() string
	Filename() string
	TypeGetter() string
	TypeFlagAdder() string
	AppliesToCmd(cmd string) bool
}

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	set := viper.Get(settingsSettingName)
	if set == nil {
		return []UserSetting{
			userSetting{
				name:            "Example",
				description:     "Add your own settings here",
				settingType:     "string",
				cmdline:         "example",
				cmdlineShortcut: "e",
				defaultVal:      "hello",
				appliesTo:       []string{"root"},
			}}
	}
	results := make([]UserSetting, len(set.([]interface{})))
	for i, item := range set.([]interface{}) {
		name := ""
		description := ""
		settingType := "string"
		cmdline := ""
		cmdlineShortcut := ""
		defaultVal := ""
		appliesTo := make([]string, 0)
		for k, v := range item.(map[string]interface{}) {
			val := fmt.Sprintf("%s", v)
			switch k {
			case "name":
				name = val
			case "description":
				description = val
			case "type":
				settingType = val
			case "cmdline":
				cmdline = val
			case "cmdlineShortcut":
				cmdlineShortcut = string(val[0])
			case "defaultVal":
				defaultVal = val
			case "appliesTo":
				for _, c := range v.([]interface{}) {
					appliesTo = append(appliesTo, c.(string))
				}
			}
		}
		results[i] = userSetting{
			name:            name,
			description:     description,
			settingType:     settingType,
			cmdline:         cmdline,
			cmdlineShortcut: cmdlineShortcut,
			defaultVal:      defaultVal,
			appliesTo:       appliesTo,
		}
	}

	return results
}

type userSetting struct {
	name            string
	description     string
	settingType     string
	cmdline         string
	cmdlineShortcut string
	defaultVal      string
	appliesTo       []string
}

func (u userSetting) Name() string {
	return strcase.ToCamel(u.name)
}

func (u userSetting) Description() string {
	return u.description
}

func (u userSetting) Type() string {
	return u.settingType
}

func (u userSetting) Cmdline() string {
	return strcase.ToKebab(u.cmdline)
}

func (u userSetting) CmdlineShortcut() string {
	return u.cmdlineShortcut
}

func (u userSetting) DefaultVal() string {
	return u.defaultVal
}

func (u userSetting) AppliesTo() []string {
	return u.appliesTo
}

func (u userSetting) Id() string {
	return strcase.ToLowerCamel(u.name)
}

func (u userSetting) Filename() string {
	return strcase.ToSnake(u.name)
}

func (u userSetting) TypeGetter() string {
	switch u.settingType {
	case "string":
		return "viper.GetString"
	case "[]string":
		return "viperEx.GetStringSlice"
	case "bool":
		return "viper.GetBool"
	case "int":
		return "viper.GetInt"
	}
	panic(errors.New(fmt.Sprintf("type '%s' is not supported as a setting type", u.settingType)))
}

func (u userSetting) TypeFlagAdder() string {
	switch u.settingType {
	case "string":
		return "viperEx.AddStringSetting"
	case "[]string":
		return "viperEx.AddStringArraySetting"
	case "bool":
		return "viperEx.AddBoolSetting"
	case "int":
		return "viperEx.AddIntSetting"
	}

	panic(errors.New(fmt.Sprintf("type '%s' is not supported as a setting type", u.settingType)))
}

func (u userSetting) AppliesToCmd(cmd string) bool {
	// No commandline, do not apply to any command
	if u.cmdline == "" {
		return false
	}
	for _, c := range u.appliesTo {
		if c == cmd {
			return true
		}
	}
	return false
}

// For template
func (u userSetting) HasPrefix(text string, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}