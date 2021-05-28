// Generated 2021-05-24 17:41:23 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

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
			UserSetting{
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
		emptyTCValues := make([]map[string]string, 0)
		results[i].optionTestCaseValues = &emptyTCValues
		results[i].appliesToRootOnly()
	}

	return results
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

var optionTestCase = template.Must(template.New(``).Parse(`"{{.OptName}}": New{{.OptType}}Option("{{.CmdLine}}"{{.Generator}}{{.Modifier}})`))

func (u UserSetting) getOptionTestCases() []map[string]string {
	if len(*u.optionTestCaseValues) == 0 {
		values := make(map[string]string)
		values["OptName"] = "Default"
		values["NameCode"] = u.NameCode()
		// Opt type / Modifier
		if strings.HasPrefix(u.Type, "[]") {
			values["OptType"] = "Slice"
			values["Modifier"] = fmt.Sprintf(`, func(s *MockSettings, value interface{}) { core.ConvertSlice(value, &s.%sVar)}`, u.NameCode())
		} else if u.Type == "bool" {
			values["OptType"] = "Boolean"
			values["Modifier"] = fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = true}`, u.NameCode())
		} else {
			values["OptType"] = "Simple"
			values["Modifier"] = fmt.Sprintf(`, func(s *MockSettings, value interface{}) { s.%sVar = value.(%s)}`, u.NameCode(), u.Type)
		}
		// Generator
		if u.Type == "bool" {
			values["Generator"] = ""
		} else {
			values["Generator"] = fmt.Sprintf(`, func() interface{} {var value %s; rg.Value(&value); return value }`, u.Type)
		}
		*u.optionTestCaseValues = append(*u.optionTestCaseValues, values)
		options := make([]map[string]string, 0, 2)
		if u.Cmdline != "" {
			options = append(options, map[string]string{
				"OptName": "Default",
				"CmdLine": fmt.Sprintf("--%s", u.Cmdline),
			})
			if u.Type == "bool" {
				options = append(options, map[string]string{
					"OptName":  "=True",
					"CmdLine":  fmt.Sprintf("--%s=true", u.Cmdline),
					"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = true}`, u.NameCode()),
				})
				options = append(options, map[string]string{
					"OptName":  "=False",
					"CmdLine":  fmt.Sprintf("--%s=false", u.Cmdline),
					"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = false}`, u.NameCode()),
				})
			}
		}
		if u.CmdlineShortcut != "" {
			options = append(options, map[string]string{
				"OptName": "Short",
				"CmdLine": fmt.Sprintf("-%s", u.CmdlineShortcut),
			})
			if u.Type == "bool" {
				options = append(options, map[string]string{
					"OptName":  "Short=True",
					"CmdLine":  fmt.Sprintf("-%s=true", u.CmdlineShortcut),
					"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = true}`, u.NameCode()),
				})
				options = append(options, map[string]string{
					"OptName":  "Short=False",
					"CmdLine":  fmt.Sprintf("-%s=false", u.CmdlineShortcut),
					"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = false}`, u.NameCode()),
				})
			}
		}
		*u.optionTestCaseValues = mapCombinations(*u.optionTestCaseValues, options)
	}
	return *u.optionTestCaseValues
}

func (u UserSetting) OptionTestCases() []string {
	lines := make([]string, 0)
	valuesList := u.getOptionTestCases()
	for _, values := range valuesList {
		buffer := bytes.Buffer{}
		optionTestCase.Execute(&buffer, values)
		lines = append(lines, buffer.String())
	}

	return lines
}

func (u UserSetting) OptionTestCaseNames() []string {
	valuesList := u.getOptionTestCases()
	names := make([]string, 0)
	for _, set := range valuesList {
		names = append(names, set["OptName"])
	}
	return names
}

func mapCombinations(existing []map[string]string, additions []map[string]string) []map[string]string {
	lenE := len(existing)
	lenA := len(additions)
	newMaps := make([]map[string]string, 0, lenA*lenE)
	for _, exist := range existing {
		for _, add := range additions {
			newMap := make(map[string]string)
			for k, v := range exist {
				newMap[k] = v
			}
			for k, v := range add {
				newMap[k] = v
			}
			newMaps = append(newMaps, newMap)
		}
	}
	return newMaps
}
