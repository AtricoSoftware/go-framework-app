package settings

type Settings interface {
	// Target directory (where to generate the app)
	TargetDirectory() string
	// Tile of application (used in readme)
	ApplicationTitle() string
	// Name of application (output will be name[.exe])
	ApplicationName() string
	// Summary of application
	ApplicationSummary() string
	// Description of application
	ApplicationDescription() string
	// Path to the repository
	RepositoryPath() string
	// Commands to add
	Commands() []UserCommand
	// Settings to add
	UserSettings() []UserSetting
	// Libraries to get
	Libraries() map[string]string
}

// Get the settings for this run
func GetSettings() Settings {
	return theSettings{}
}

// Stub object for settings interface
type theSettings struct{}
