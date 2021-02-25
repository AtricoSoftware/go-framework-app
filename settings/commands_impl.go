// Generated 2021-02-24 16:31:35 by go-framework v1.5.0
package settings

import (
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

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
