package resources

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/atrico-go/display"
)

var FuncMap = template.FuncMap{
	// Concatenate slices
	"concat": func(lists ...[]string) (result []string) {
		for _, lst := range lists {
			result = append(result, lst...)
		}
		return result
	},
	// Comma separated list
	"commaList": func(lst []string) string { return strings.Join(lst, ", ") },
	// Quote items in a list
	"quoted": func(lst []string) (result []string) {
		result = make([]string, len(lst))
		for i, item := range lst {
			result[i] = fmt.Sprintf(`"%s"`, item)
		}
		return result
	},
	// Create a table for text alignment
	"createTable": func() display.TableBuilder {
		return display.NewTableBuilder().WithVerticalSeparator(' ')
	},
	// Display the table created above
	"printTable": func(table display.TableBuilder) string {
		text := strings.Builder{}
		for _,line := range table.Build().Render().StringMl() {
			text.WriteString(strings.TrimRight(line, " "))
			text.WriteString("\n")
		}
		return text.String()
	},
}
