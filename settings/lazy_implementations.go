// Generated 2021-03-09 17:48:01 by go-framework V1.8.0
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

// ----------------------------------------------------------------------------------------------------------------------------
// Bool
// ----------------------------------------------------------------------------------------------------------------------------
type lazyBoolValue lazyValue

type LazyBoolValue interface {
	GetValue() bool
}

func NewLazyBoolValue(creator func() bool) LazyBoolValue {
	lz := lazyBoolValue(lazyValue{creator: func() interface{} { return creator() }})
	return &lz
}

func (v *lazyBoolValue) GetValue() bool {
	return ((*lazyValue)(v).getValue()).(bool)
}

// ----------------------------------------------------------------------------------------------------------------------------
// String
// ----------------------------------------------------------------------------------------------------------------------------
type lazyStringValue lazyValue

type LazyStringValue interface {
	GetValue() string
}

func NewLazyStringValue(creator func() string) LazyStringValue {
	lz := lazyStringValue(lazyValue{creator: func() interface{} { return creator() }})
	return &lz
}

func (v *lazyStringValue) GetValue() string {
	return ((*lazyValue)(v).getValue()).(string)
}

// ----------------------------------------------------------------------------------------------------------------------------
// UserCommandSlice
// ----------------------------------------------------------------------------------------------------------------------------
type lazyUserCommandSliceValue lazyValue

type LazyUserCommandSliceValue interface {
	GetValue() []UserCommand
}

func NewLazyUserCommandSliceValue(creator func() []UserCommand) LazyUserCommandSliceValue {
	lz := lazyUserCommandSliceValue(lazyValue{creator: func() interface{} { return creator() }})
	return &lz
}

func (v *lazyUserCommandSliceValue) GetValue() []UserCommand {
	return ((*lazyValue)(v).getValue()).([]UserCommand)
}

// ----------------------------------------------------------------------------------------------------------------------------
// UserSettingSlice
// ----------------------------------------------------------------------------------------------------------------------------
type lazyUserSettingSliceValue lazyValue

type LazyUserSettingSliceValue interface {
	GetValue() []UserSetting
}

func NewLazyUserSettingSliceValue(creator func() []UserSetting) LazyUserSettingSliceValue {
	lz := lazyUserSettingSliceValue(lazyValue{creator: func() interface{} { return creator() }})
	return &lz
}

func (v *lazyUserSettingSliceValue) GetValue() []UserSetting {
	return ((*lazyValue)(v).getValue()).([]UserSetting)
}
