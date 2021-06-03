// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
package unit_tests

import (
	"fmt"
	"strings"

	"github.com/atrico-go/core"
)

// ----------------------------------------------------------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------------------------------------------------------

type Option interface {
	Set()
	Cmdline() string
	ModifySettings(settings *MockSettings)
}
type OptionSet map[string]Option
type OptionSetList []OptionSet

// ----------------------------------------------------------------------------------------------------------------------------
// Simple option (--opt <val>)
// ----------------------------------------------------------------------------------------------------------------------------

func NewSimpleOption(option string, generator func() interface{}, modifier func(s *MockSettings, value interface{})) Option {
	opt := simpleOption{
		option:    option,
		generator: generator,
		modifier:  modifier,
	}
	return &opt
}

type simpleOption struct {
	value     interface{}
	option    string
	generator func() interface{}
	modifier  func(s *MockSettings, value interface{})
}

func (o *simpleOption) Set() {
	o.value = o.generator()
}

func (o *simpleOption) Cmdline() string {
	return fmt.Sprintf("%s %v", o.option, o.value)
}

func (o *simpleOption) ModifySettings(settings *MockSettings) {
	o.modifier(settings, o.value)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Boolean option (--opt/--opt=true/--opt=false)
// ----------------------------------------------------------------------------------------------------------------------------

func NewBooleanOption(option string, modifier func(s *MockSettings)) Option {
	opt := boolOption{
		option:   option,
		modifier: modifier,
	}
	return &opt
}

type boolOption struct {
	value    bool
	option   string
	modifier func(s *MockSettings)
}

func (o *boolOption) Set() {
	// Nothing to do
}

func (o *boolOption) Cmdline() string {
	return fmt.Sprintf("%s", o.option)
}

func (o *boolOption) ModifySettings(settings *MockSettings) {
	o.modifier(settings)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Slice option (--opt 1 --opt 2 --opt 3)
// ----------------------------------------------------------------------------------------------------------------------------

func NewSliceOption(option string, generator func() interface{}, modifier func(s *MockSettings, value interface{})) Option {
	opt := sliceOption{
		option:    option,
		generator: generator,
		modifier:  modifier,
	}
	return &opt
}

type sliceOption struct {
	value     []interface{}
	option    string
	generator func() interface{}
	modifier  func(s *MockSettings, value interface{})
}

func (o *sliceOption) Set() {
	core.ConvertSlice(o.generator(), &o.value)
}

func (o *sliceOption) Cmdline() string {
	cmdline := strings.Builder{}
	for _, item := range o.value {
		cmdline.WriteString(fmt.Sprintf("%s %v ", o.option, item))
	}
	return cmdline.String()
}

func (o *sliceOption) ModifySettings(settings *MockSettings) {
	o.modifier(settings, o.value)
}
