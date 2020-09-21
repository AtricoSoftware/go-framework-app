package settings

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

const commandsSettingName = "commands"

type UserCommand interface {
	Name() string
	Description() string

	ApiName() string
	UseName() string
}

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	set := viper.Get(commandsSettingName)
	if set == nil {
		return []UserCommand{
			userCommand{
				name:        "example",
				description: "Example command",
			}}
	}
	results := make([]UserCommand, len(set.([]interface{})))
	for i, item := range set.([]interface{}) {
		name := ""
		description := ""
		for k, v := range item.(map[string]interface{}) {
			val := fmt.Sprintf("%s", v)
			switch k {
			case "name":
				name = val
			case "description":
				description = val
			}
		}
		results[i] = userCommand{
			name:        name,
			description: description,
		}
	}

	return results
}

type userCommand struct {
	name        string
	description string
}

func (c userCommand) Name() string        { return strcase.ToLowerCamel(c.name) }
func (c userCommand) Description() string { return c.description }
func (c userCommand) ApiName() string     { return strcase.ToCamel(c.name) }
func (c userCommand) UseName() string     { return strcase.ToKebab(c.name) }
