package file_writer

import (
	"reflect"

	"github.com/AtricoSoftware/go-framework-app/settings"
)

type TemplateValues map[string]interface{}

func CreateTemplateValues(settings settings.Settings) TemplateValues {
	t := reflect.TypeOf(settings)
	numMethods := t.NumMethod()
	values := make(TemplateValues, numMethods)
	for i := 0; i < numMethods; i++ {
		method := t.Method(i)
		values[method.Name] = reflect.ValueOf(settings).MethodByName(method.Name).Call([]reflect.Value{})[0].Interface()
	}
	return values
}

