// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
package unit_tests

import (
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type MockSettings struct {
	TheCommand                 []string
	TheArgs                    []string
	SingleReadConfigurationVar bool
	IncludeDryRunVar           bool
	TargetDirectoryVar         string
	ApplicationTitleVar        string
	ApplicationNameVar         string
	ApplicationSummaryVar      string
	ApplicationDescriptionVar  string
	RepositoryPathVar          string
	CommandsVar                []settings.UserCommand
	UserSettingsVar            []settings.UserSetting
}

func (s MockSettings) SingleReadConfiguration() bool {
	return s.SingleReadConfigurationVar
}
func (s MockSettings) IncludeDryRun() bool {
	return s.IncludeDryRunVar
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

func NewMockSettings(cmd []string, args []string) MockSettings {
	return MockSettings{
		TheCommand:                 cmd,
		TheArgs:                    args,
		SingleReadConfigurationVar: true,
		IncludeDryRunVar:           true,
		TargetDirectoryVar:         ".",
		CommandsVar:                make([]settings.UserCommand, 0),
		UserSettingsVar:            make([]settings.UserSetting, 0),
	}
}
