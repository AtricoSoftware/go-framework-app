// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
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

// Stub object for settings interface
type theSettings struct{}
