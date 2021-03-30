// Generated 2021-03-30 15:32:41 by go-framework development-version
package unit_tests

type MockSettings struct {
	TheCommand	string
	SingleReadConfigurationVar bool
	TargetDirectoryVar string
	ApplicationTitleVar string
	ApplicationNameVar string
	ApplicationSummaryVar string
	ApplicationDescriptionVar string
	RepositoryPathVar string
	CommandsVar []settings.UserCommand
	UserSettingsVar []settings.UserSetting
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

func NewMockSettings(cmd string) MockSettings {
	return MockSettings{
		TheCommand: cmd,
		CommandsVar: make([]settings.UserCommand, 0),
		UserSettingsVar: make([]settings.UserSetting, 0),
	}
}
