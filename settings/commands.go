// Generated 2021-03-30 15:32:41 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/viper"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

const commandsSettingName = "Commands"

// Cached value
var commandsSettingCache = NewCachedUserCommandSliceValue(func() []UserCommand { return ParseCommandsSetting(viper.Get(commandsSettingName)) })

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	return commandsSettingCache.GetValue()
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
