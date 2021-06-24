package settings

type Settings interface {
	// Target directory (where to generate the app)
	TargetDirectory() string
	// Name of application (output will be name[.exe])
	ApplicationName() string
	// Path to the repository
	RepositoryPath() string
	// Commands to add
	Commands() []string
	// Settings to add
	UserSettings() []UserSetting
}

// Get the settings for this run
func GetSettings() Settings {
	return theSettings{}
}

// Stub object for settings interface
type theSettings struct{}
