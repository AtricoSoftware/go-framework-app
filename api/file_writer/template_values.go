package file_writer

import (
	"reflect"
)

type TemplateValues map[string]interface{}

func (f fileWriter) CreateTemplateValues() TemplateValues {
	t := reflect.TypeOf(f.config)
	numMethods := t.NumMethod()
	values := make(TemplateValues, numMethods)
	for i := 0; i < numMethods; i++ {
		method := t.Method(i)
		// Skip argument functions
		if method.Name != "GetArgument" && method.Name != "MustGetArgument" && method.Name != "SetArgs" {
			values[method.Name] = reflect.ValueOf(f.config).MethodByName(method.Name).Call([]reflect.Value{})[0].Interface()
		}
	}
	// Add comment string/backup suffix to values
	values["Comment"] = f.fileComment()
	values["BackupSuffix"] = f.now.Format("2006-01-02_15-04-05")
	return values
}
