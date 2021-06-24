// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
package unit_tests

import (
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type MockSettings struct {
	TheCommand                 []string
	TheArgs                    []string
	SingleReadConfigurationVar bool
	TargetDirectoryVar         string
	ApplicationTitleVar        string
	ApplicationNameVar         string
	ApplicationSummaryVar      string
	ApplicationDescriptionVar  string
	RepositoryPathVar          string
	CommandsVar                []settings.UserCommand
	UserSettingsVar            []settings.UserSetting
	SkeletonFilesVar           []string
	ConfigFileVar              string
	VerboseVar                 bool
}

func (s MockSettings) SingleReadConfiguration() bool {
	return s.SingleReadConfigurationVar
}
func (s MockSettings) TargetDirectory() string {
	return s.TargetDirectoryVar
}
func (s MockSettings) ApplicationTitle() string {
	return s.ApplicationTitleVar
}
func (s MockSettings) ApplicationName() string {
	return s.ApplicationNameVar
}
func (s MockSettings) ApplicationSummary() string {
	return s.ApplicationSummaryVar
}
func (s MockSettings) ApplicationDescription() string {
	return s.ApplicationDescriptionVar
}
func (s MockSettings) RepositoryPath() string {
	return s.RepositoryPathVar
}
func (s MockSettings) Commands() []settings.UserCommand {
	return s.CommandsVar
}
func (s MockSettings) UserSettings() []settings.UserSetting {
	return s.UserSettingsVar
}
func (s MockSettings) SkeletonFiles() []string {
	return s.SkeletonFilesVar
}
func (s MockSettings) ConfigFile() string {
	return s.ConfigFileVar
}
func (s MockSettings) Verbose() bool {
	return s.VerboseVar
}

func NewMockSettings(cmd []string, args []string) MockSettings {
	return MockSettings{
		TheCommand:                 cmd,
		TheArgs:                    args,
		SingleReadConfigurationVar: true,
		TargetDirectoryVar:         ".",
		CommandsVar:                make([]settings.UserCommand, 0),
		UserSettingsVar:            make([]settings.UserSetting, 0),
		SkeletonFilesVar:           make([]string, 0),
	}
}
