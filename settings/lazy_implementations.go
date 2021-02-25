package settings

type lazyValue struct {
	theValue interface{}
	creator func() interface{}
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
// String
// ----------------------------------------------------------------------------------------------------------------------------

type LazyStringValue interface {
	GetValue() string
}
type lazyStringValue lazyValue
func NewLazyStringValue(creator func() string) LazyStringValue {
	lz := lazyStringValue(lazyValue{creator:func() interface{} {return creator()}})
	return &lz
}
func (v *lazyStringValue) GetValue() string {
	return ((*lazyValue)(v).getValue()).(string)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Bool
// ----------------------------------------------------------------------------------------------------------------------------

type LazyBoolValue interface {
	GetValue() bool
}
type lazyBoolValue lazyValue
func NewLazyBoolValue(creator func() bool) LazyBoolValue {
	lz := lazyBoolValue(lazyValue{creator:func() interface{} {return creator()}})
	return &lz
}
func (v *lazyBoolValue) GetValue() bool {
	return ((*lazyValue)(v).getValue()).(bool)
}

// ----------------------------------------------------------------------------------------------------------------------------
// String Slice
// ----------------------------------------------------------------------------------------------------------------------------

type LazyStringSliceValue interface {
	GetValue() []string
}
type lazyStringSliceValue lazyValue
func NewLazyStringSliceValue(creator func() []string) LazyStringSliceValue {
	lz := lazyStringSliceValue(lazyValue{creator:func() interface{} {return creator()}})
	return &lz
}
func (v *lazyStringSliceValue) GetValue() []string {
	return ((*lazyValue)(v).getValue()).([]string)
}

// ----------------------------------------------------------------------------------------------------------------------------
// UserCommand Slice
// ----------------------------------------------------------------------------------------------------------------------------

type LazyUserCommandSliceValue interface {
	GetValue() []UserCommand
}
type lazyUserCommandSliceValue lazyValue
func NewLazyUserCommandSliceValue(creator func() []UserCommand) LazyUserCommandSliceValue {
	lz := lazyUserCommandSliceValue(lazyValue{creator:func() interface{} {return creator()}})
	return &lz
}
func (v *lazyUserCommandSliceValue) GetValue() []UserCommand {
	return ((*lazyValue)(v).getValue()).([]UserCommand)
}

// ----------------------------------------------------------------------------------------------------------------------------
// UserSetting Slice
// ----------------------------------------------------------------------------------------------------------------------------

type LazyUserSettingSliceValue interface {
	GetValue() []UserSetting
}
type lazyUserSettingSliceValue lazyValue
func NewLazyUserSettingSliceValue(creator func() []UserSetting) LazyUserSettingSliceValue {
	lz := lazyUserSettingSliceValue(lazyValue{creator:func() interface{} {return creator()}})
	return &lz
}
func (v *lazyUserSettingSliceValue) GetValue() []UserSetting {
	return ((*lazyValue)(v).getValue()).([]UserSetting)
}
