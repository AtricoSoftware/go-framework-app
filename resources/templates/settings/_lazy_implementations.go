// {{.Comment}}
package settings

type lazyValue struct {
	theValue interface{}
	creator  func() interface{}
	hasValue bool
}

func (v *lazyValue) getValue() interface{} {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}
{{- range .LazySettings}}

// ----------------------------------------------------------------------------------------------------------------------------
// {{.TypeNameAsCode}}
// ----------------------------------------------------------------------------------------------------------------------------
{{- $lazyValue := print "Lazy" .TypeNameAsCode "Value"}}
{{- $lazyValueLow := print "lazy" .TypeNameAsCode "Value"}}
type {{$lazyValueLow}} lazyValue

type {{$lazyValue}} interface {
	GetValue() {{.Type}}
}

func New{{$lazyValue}}(creator func() {{.Type}}) {{$lazyValue}} {
	lz := {{$lazyValueLow}}(lazyValue{creator: func() interface{} { return creator() }})
	return &lz
}

func (v *{{$lazyValueLow}}) GetValue() {{.Type}} {
	return ((*lazyValue)(v).getValue()).({{.Type}})
}
{{- end}}
