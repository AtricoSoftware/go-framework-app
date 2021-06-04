// Generated 2021-06-04 15:53:11 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/viper"
	"fmt"
	"path"
	"regexp"
	"strings"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

const commandsSettingName = "Commands"

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	return ParseCommandsSetting(viper.Get(commandsSettingName))
}
// SECTION-END

type UserCommand struct {
	Name             string
	Description      string
	NoImplementation bool
	Args             []string
	OptionalArgs     []string
}

func ParseCommandsSetting(setting interface{}) []UserCommand {
	if setting == nil {
		return []UserCommand{
			{
				Name:        "example",
				Description: "Example command",
				Args:        []string{"param1"},
			}}
	}
	results := make([]UserCommand, len(setting.([]interface{})))
	for i, item := range setting.([]interface{}) {
		mapstructure.Decode(item, &(results[i]))
	}
	return results
}

func (c UserCommand) ApiName() string      { return strcase.ToCamel(c.stripPath()) }
func (c UserCommand) LowerApiName() string { return strcase.ToLowerCamel(c.stripPath()) }
func (c UserCommand) UseName() string      { return ToKebabEx(path.Base(c.Name)) }
func (c UserCommand) FileName() string     { return ToKebabEx(c.stripPath()) }
func (c UserCommand) HasArgs() bool        { return len(c.Args)+len(c.OptionalArgs) > 0 }
func (c UserCommand) SplitPath() []string  { return strings.Split(ToKebabEx(c.Name), "/") }
func (c UserCommand) ArgsValidator() string {
	if len(c.Args) == 0 {
		if len(c.OptionalArgs) == 0 {
			return "NoArgs"
		} else {
			return fmt.Sprintf("MaximumNArgs(%d)", len(c.OptionalArgs))
		}
	} else {
		if len(c.OptionalArgs) == 0 {
			return fmt.Sprintf("ExactArgs(%d)", len(c.Args))
		} else {
			return fmt.Sprintf("RangeArgs(%d, %d)", len(c.Args), len(c.Args)+len(c.OptionalArgs))
		}
	}
}
func (c UserCommand) stripPath() string { return strings.ReplaceAll(c.Name, "/", " ") }

var regExShish = regexp.MustCompile(`\-(\d+)`)

func ToKebabEx(input string) string {
	str := strcase.ToKebab(input)
	// Strip out delimiters before numbers
	return regExShish.ReplaceAllString(str, "$1")
}
