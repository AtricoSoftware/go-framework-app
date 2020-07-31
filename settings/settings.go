package settings

type Settings interface {
	// TODO - Add your own settings as required
	Example() string
}

// Get the settings for this run
func GetSettings() Settings {
	return theSettings{}
}

// Stub object for settings interface
type theSettings struct{}
