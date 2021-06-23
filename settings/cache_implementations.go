// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
package settings

// ----------------------------------------------------------------------------------------------------------------------------
// Bool
// ----------------------------------------------------------------------------------------------------------------------------
type CachedBoolValue interface {
	GetValue() bool
	Reset()
}

func NewCachedBoolValue(creator func() bool) CachedBoolValue {
	cv := cachedBoolValue{creator: creator}
	return &cv
}

type cachedBoolValue struct {
	theValue bool
	creator  func() bool
	hasValue bool
}

func (v *cachedBoolValue) GetValue() bool {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}

func (v *cachedBoolValue) Reset() {
	v.hasValue = false
}

// ----------------------------------------------------------------------------------------------------------------------------
// String
// ----------------------------------------------------------------------------------------------------------------------------
type CachedStringValue interface {
	GetValue() string
	Reset()
}

func NewCachedStringValue(creator func() string) CachedStringValue {
	cv := cachedStringValue{creator: creator}
	return &cv
}

type cachedStringValue struct {
	theValue string
	creator  func() string
	hasValue bool
}

func (v *cachedStringValue) GetValue() string {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}

func (v *cachedStringValue) Reset() {
	v.hasValue = false
}

// ----------------------------------------------------------------------------------------------------------------------------
// StringSlice
// ----------------------------------------------------------------------------------------------------------------------------
type CachedStringSliceValue interface {
	GetValue() []string
	Reset()
}

func NewCachedStringSliceValue(creator func() []string) CachedStringSliceValue {
	cv := cachedStringSliceValue{creator: creator}
	return &cv
}

type cachedStringSliceValue struct {
	theValue []string
	creator  func() []string
	hasValue bool
}

func (v *cachedStringSliceValue) GetValue() []string {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}

func (v *cachedStringSliceValue) Reset() {
	v.hasValue = false
}

// ----------------------------------------------------------------------------------------------------------------------------
// UserCommandSlice
// ----------------------------------------------------------------------------------------------------------------------------
type CachedUserCommandSliceValue interface {
	GetValue() []UserCommand
	Reset()
}

func NewCachedUserCommandSliceValue(creator func() []UserCommand) CachedUserCommandSliceValue {
	cv := cachedUserCommandSliceValue{creator: creator}
	return &cv
}

type cachedUserCommandSliceValue struct {
	theValue []UserCommand
	creator  func() []UserCommand
	hasValue bool
}

func (v *cachedUserCommandSliceValue) GetValue() []UserCommand {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}

func (v *cachedUserCommandSliceValue) Reset() {
	v.hasValue = false
}

// ----------------------------------------------------------------------------------------------------------------------------
// UserSettingSlice
// ----------------------------------------------------------------------------------------------------------------------------
type CachedUserSettingSliceValue interface {
	GetValue() []UserSetting
	Reset()
}

func NewCachedUserSettingSliceValue(creator func() []UserSetting) CachedUserSettingSliceValue {
	cv := cachedUserSettingSliceValue{creator: creator}
	return &cv
}

type cachedUserSettingSliceValue struct {
	theValue []UserSetting
	creator  func() []UserSetting
	hasValue bool
}

func (v *cachedUserSettingSliceValue) GetValue() []UserSetting {
	if !v.hasValue {
		v.theValue = v.creator()
		v.hasValue = true
	}
	return v.theValue
}

func (v *cachedUserSettingSliceValue) Reset() {
	v.hasValue = false
}
