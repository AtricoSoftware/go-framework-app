package settings

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var optionTestCase = template.Must(template.New(``).Parse(`"{{.OptName}}": New{{.OptType}}Option("{{.CmdLine}}"{{.Generator}}{{.Modifier}})`))

func (u UserSetting) getOptionTestCases() []map[string]string {
	// No tests for config-file (this will fail before test reaches mock)
	if u.Name != "ConfigFile" {
		if len(*u.optionTestCaseValues) == 0 {
			values := make(map[string]string)
			values["OptName"] = "Default"
			values["NameCode"] = u.NameCode()
			// Opt type / Modifier
			if strings.HasPrefix(u.Type, "[]") {
				values["OptType"] = "Slice"
				values["Modifier"] = fmt.Sprintf(`, func(s *MockSettings, value interface{}) { core.ConvertSlice(value, &s.%sVar) }`, u.NameCode())
			} else if u.Type == "bool" {
				values["OptType"] = "Boolean"
				values["Modifier"] = fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = true }`, u.NameCode())
			} else {
				values["OptType"] = "Simple"
				values["Modifier"] = fmt.Sprintf(`, func(s *MockSettings, value interface{}) { s.%sVar = value.(%s) }`, u.NameCode(), u.Type)
			}
			// Generator
			if u.Type == "bool" {
				values["Generator"] = ""
			} else {
				values["Generator"] = fmt.Sprintf(`, func() interface{} { var value %s; rg.Value(&value); return value }`, u.Type)
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
						"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = true }`, u.NameCode()),
					})
					options = append(options, map[string]string{
						"OptName":  "=False",
						"CmdLine":  fmt.Sprintf("--%s=false", u.Cmdline),
						"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = false }`, u.NameCode()),
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
						"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = true }`, u.NameCode()),
					})
					options = append(options, map[string]string{
						"OptName":  "Short=False",
						"CmdLine":  fmt.Sprintf("-%s=false", u.CmdlineShortcut),
						"Modifier": fmt.Sprintf(`, func(s *MockSettings) { s.%sVar = false }`, u.NameCode()),
					})
				}
			}
			*u.optionTestCaseValues = mapCombinations(*u.optionTestCaseValues, options)
		}
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
