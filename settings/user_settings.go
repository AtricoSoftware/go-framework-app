// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
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
	"string":   {typeGetter: "viper.GetString", typeFlagAdder: "viperEx.AddStringSetting"},
	"[]string": {typeGetter: "viperEx.GetStringSlice", typeFlagAdder: "viperEx.AddStringArraySetting"},
	"bool":     {typeGetter: "viper.GetBool", typeFlagAdder: "viperEx.AddBoolSetting"},
	"int":      {typeGetter: "viper.GetInt", typeFlagAdder: "viperEx.AddIntSetting"},
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

// ----------------------------------------------------------------------------------------------------------------------------
// Delete below this
// ----------------------------------------------------------------------------------------------------------------------------

// // Return appropriate test data
// // In a form that will compile as code
// func (u UserSetting) GetTestDataCode() string {
// 	return u.testDataCode
// }
//
// // Return appropriate test data
// // In a form to use on cmdline
// func (u UserSetting) GetTestDataCmdline() string {
// 	return u.testDataCmdline
// }
//
// func (u *UserSetting) GenerateTestData(cmdline string) string {
// 	switch u.Type {
// 	case "string":
// 		val := randomString()
// 		u.testDataCmdline = fmt.Sprintf(`%s %s`, cmdline, val)
// 		u.testDataCode = fmt.Sprintf(`"%s"`, val)
// 	case "[]string":
// 		val := randomStringSlice()
// 		cmdlineSep := fmt.Sprintf(` %s `, cmdline)
// 		cmdlineFormat := fmt.Sprintf("%s %%s", cmdline)
// 		u.testDataCmdline = fmt.Sprintf(cmdlineFormat, strings.Join(val, cmdlineSep)) // Relies on slice having at least 1 entry
// 		u.testDataCode = fmt.Sprintf(`[]string{"%s"}`, strings.Join(val, `","`))
// 	case "bool":
// 		// TODO - false as possibility
// 		u.testDataCmdline = cmdline
// 		u.testDataCode = "true"
// 	case "int":
// 		val := rand.Intn(1000)
// 		u.testDataCmdline = fmt.Sprintf(`%s %d`, cmdline, val)
// 		u.testDataCode = fmt.Sprintf(`%d`, val)
// 	}
// 	return ""
// }
//
// // ----------------------------------------------------------------------------------------------------------------------------
// // Random values
// // ----------------------------------------------------------------------------------------------------------------------------
//
// func randomRune() rune {
// 	val := rand.Intn(62)
// 	switch {
// 	case val < 26:
// 		return rune('A' + val)
// 	case val < 52:
// 		return rune('a' + val - 26)
// 	default:
// 		return rune('0' + val - 52)
// 	}
// }
//
// func randomString() string {
// 	str := strings.Builder{}
// 	for i := 0; i < 5; i++ {
// 		str.WriteRune(randomRune())
// 	}
// 	return str.String()
// }
//
// func randomStringSlice() []string {
// 	slice := make([]string, 3)
// 	for i := 0; i < 3; i++ {
// 		slice[i] = randomString()
// 	}
// 	return slice
// }
