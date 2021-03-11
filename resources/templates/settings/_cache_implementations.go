// {{.Comment}}
package settings
{{- range .CachedTypes}}

// ----------------------------------------------------------------------------------------------------------------------------
// {{.TypeNameAsCode}}
// ----------------------------------------------------------------------------------------------------------------------------
{{- $interfaceName := print "Cached" .TypeNameAsCode "Value"}}
{{- $structName := print "cached" .TypeNameAsCode "Value"}}
type {{$interfaceName}} interface {
	GetValue() {{.Type}}
	Reset()
}

func New{{$interfaceName}}(creator func() {{.Type}}) {{$interfaceName}} {
	cv := {{$structName}}{creator: creator }
	return &cv
}

type {{$structName}} struct {
	theValue {{.Type}}
	creator  func() {{.Type}}
	hasValue bool
}

func (v *{{$structName}}) GetValue() {{.Type}} {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}

func (v *{{$structName}}) Reset() {
	v.hasValue = false
}
{{- end}}
