// Generated 2021-02-25 13:40:05 by go-framework development-version
package settings

import "github.com/atrico-go/container"

type Settings interface {
	// Configuration is only read once (at start)
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
	// Libraries to get
	Libraries() []string
}

// Register the settings
func RegisterSettings(c container.Container) {
	c.Singleton(func() Settings { return theSettings{} })
}

// Stub object for settings interface
type theSettings struct{}
