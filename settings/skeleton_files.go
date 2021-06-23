// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
)

const skeletonFilesSettingName = "SkeletonFiles"
const skeletonFilesSettingCmdline = "skeleton"
const skeletonFilesSettingShortcut = 's'

// Cached value
var skeletonFilesSettingCache = NewCachedStringSliceValue(func() []string { return viperEx.GetStringSlice(skeletonFilesSettingName) })

// Fetch the setting
func (theSettings) SkeletonFiles() []string {
	return skeletonFilesSettingCache.GetValue()
}

func AddSkeletonFilesFlag(flagSet *pflag.FlagSet) {
	viperEx.StringArraySetting(skeletonFilesSettingName, "File(s) with skeleton definitions").Cmdline(skeletonFilesSettingCmdline).CmdlineShortcut(skeletonFilesSettingShortcut).AddTo(flagSet)
}

// SECTION-END
