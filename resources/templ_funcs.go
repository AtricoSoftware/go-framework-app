package resources

import (
	"fmt"
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	// Concatenate slices
	"concat": func(lists ...[]string) (result []string) {
		for _,lst := range lists {
			result = append(result, lst...)
		}
		return result
	},
	// Comma separated list
	"commaList": func(lst []string) string {return strings.Join(lst, ",")},
	// Quote items in a list
	"quoted": func(lst []string) (result []string) {
		result = make([]string, len(lst))
		for i,item := range lst {
			result[i] = fmt.Sprintf(`"%s"`, item)
		}
		return result
	},
}
