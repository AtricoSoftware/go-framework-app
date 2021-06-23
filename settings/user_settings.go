// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
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

// Cached value
var userSettingsSettingCache = NewCachedUserSettingSliceValue(func() []UserSetting { return ParseUserSettingsSetting(viper.Get(userSettingsSettingName)) })

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	return userSettingsSettingCache.GetValue()
}

// SECTION-END

type UserSetting struct {
	Skeleton             string
	Name                 string
	Id                   string
	Description          string
	Type                 string
	Cmdline              string
	CmdlineShortcut      string
	EnvVar               string
	DefaultVal           string
	AppliesTo            []string
	testDataCmdline      string
	testDataCode         string
	optionTestCaseValues *[]map[string]string
}

func ParseUserSettingsSetting(setting interface{}) []UserSetting {
	if setting == nil {
		emptyTCValues := make([]map[string]string, 0)
		return []UserSetting{
			{
				Name:                 "Example", // Name of setting - used in const
				Id:                   "example", // viper ID, dotted values form groups
				Description:          "Add your own settings here",
				Type:                 "string",
				Cmdline:              "example",
				CmdlineShortcut:      "e",
				EnvVar:               "EXAMPLE",
				DefaultVal:           "hello",
				AppliesTo:            []string{"root"},
				optionTestCaseValues: &emptyTCValues,
			}}
	}
	results := make([]UserSetting, len(setting.([]interface{})))
	for i, item := range setting.([]interface{}) {
		mapstructure.Decode(item, &(results[i]))
		results[i] = resolveSkeleton(results[i])
		results[i].optionTestCaseValues = emptyTcValues()
		results[i].appliesToRootOnly()
	}
	// Add extra settings
	results = append(results, generateExtraSettings()...)
	return results
}

func emptyTcValues() *[]map[string]string {
	emptyTCValues := make([]map[string]string, 0)
	return &emptyTCValues
}

type typeDetails struct {
	typeGetter    string
	typeFlagAdder string
}

var typeInfo = map[string]typeDetails{
	"string":   {typeGetter: "viper.GetString", typeFlagAdder: "viperEx.StringSetting"},
	"[]string": {typeGetter: "viperEx.GetStringSlice", typeFlagAdder: "viperEx.StringArraySetting"},
	"bool":     {typeGetter: "viper.GetBool", typeFlagAdder: "viperEx.BoolSetting"},
	"int":      {typeGetter: "viper.GetInt", typeFlagAdder: "viperEx.IntSetting"},
}

func (u *UserSetting) appliesToRootOnly() {
	// If nil or empty, set to root
	// If contains root, remove others
	if u.AppliesTo == nil || len(u.AppliesTo) == 0 || (len(u.AppliesTo) > 1 && sliceContains(u.AppliesTo, "root")) {
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

func (u UserSetting) RawType() string {
	raw, _ := u.rawType()
	return raw
}
func (u UserSetting) rawType() (raw, prefix string) {
	raw = u.Type
	for strings.HasPrefix(raw, "[]") {
		raw = raw[2:]
		prefix = prefix + "[]"
	}
	return raw, prefix

}

func (u UserSetting) QualifiedType() string {
	if _, ok := typeInfo[u.Type]; !ok {
		raw, prefix := u.rawType()
		return fmt.Sprintf("%ssettings.%s", prefix, raw)
	}
	return u.Type
}

func (u UserSetting) TypeGetter() string {
	if info, ok := typeInfo[u.Type]; ok {
		return info.typeGetter
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
	if info, ok := typeInfo[u.Type]; ok {
		return info.typeFlagAdder
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

func (u UserSetting) AppliesToCmdOrRoot(cmd string) bool {
	return u.AppliesTo[0] == "root" || u.AppliesToCmd(cmd)
}

// For template
func (u UserSetting) HasPrefix(text string, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}

// Init code for this type ("" if doesn't need it)
func (u UserSetting) InitCode() string {
	return initCode(u.Type)
}
func (u UserSetting) QualifiedInitCode() string {
	return initCode(u.QualifiedType())
}
func initCode(ty string) string {
	if strings.HasPrefix(ty, "[]") {
		return fmt.Sprintf(`make(%s, 0)`, ty)
	}
	if strings.HasPrefix(ty, "map") {
		return fmt.Sprintf(`make(%s)`, ty)
	}
	return ""
}

// Applies to as comma separated list
func (u UserSetting) AppliesToCSL() string {
	return strings.Replace(strings.Join(u.AppliesTo, ", "), "root", "all", 1)
}
