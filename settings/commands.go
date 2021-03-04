// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
// SECTION-START: Framework
package settings

import (
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const commandsSettingName = "Commands"

// Lazy value
var commandsSettingLazy = NewLazyUserCommandSliceValue(func() []UserCommand { return ParseCommandsSetting(viper.Get(commandsSettingName)) })

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	return commandsSettingLazy.GetValue()
}

// SECTION-END

type UserCommand struct {
	Name        string
	Description string
}

func ParseCommandsSetting(setting interface{}) []UserCommand {
	if setting == nil {
		return []UserCommand{
			{
				Name:        "example",
				Description: "Example command",
			}}
	}
	results := make([]UserCommand, len(setting.([]interface{})))
	for i, item := range setting.([]interface{}) {
		mapstructure.Decode(item, &(results[i]))
	}
	return results
}

func (c UserCommand) ApiName() string { return strcase.ToCamel(c.Name) }
func (c UserCommand) UseName() string { return strcase.ToKebab(c.Name) }
