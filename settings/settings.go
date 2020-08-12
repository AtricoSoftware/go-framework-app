package settings

type Settings interface {
	// Target directory (where to generate the app)
	TargetDirectory() string
	// Name of application (output will be name[.exe])
	ApplicationName() string
	// Path to the repository
	RepositoryPath() string
}

// Get the settings for this run
func GetSettings() Settings {
	return theSettings{}
}

// Stub object for settings interface
type theSettings struct{}
