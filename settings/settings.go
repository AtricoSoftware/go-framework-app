// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
package settings

import "github.com/atrico-go/container"

type Settings interface {
	// Configuration is only read once (at startup)
	SingleReadConfiguration() bool
	// Target directory
	TargetDirectory() string
	// Name of application
	ApplicationTitle() string
	// Name of application
	ApplicationName() string
	// Summary description of application
	ApplicationSummary() string
	// Description of application
	ApplicationDescription() string
	// Path to repository
	RepositoryPath() string
	// Commands to add
	Commands() []UserCommand
	// Settings to add
	UserSettings() []UserSetting
}

// Register the settings
func RegisterSettings(c container.Container) {
	c.Singleton(func() Settings { return theSettings{} })
}

// Force all settings to be recalculated on next request
func ResetCaches() {
	singleReadConfigurationSettingCache.Reset()
	targetDirectorySettingCache.Reset()
	applicationTitleSettingCache.Reset()
	applicationNameSettingCache.Reset()
	applicationSummarySettingCache.Reset()
	applicationDescriptionSettingCache.Reset()
	repositoryPathSettingCache.Reset()
	commandsSettingCache.Reset()
	userSettingsSettingCache.Reset()
}

// Stub object for settings interface
type theSettings struct{}
