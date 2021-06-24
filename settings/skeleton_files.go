// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
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
