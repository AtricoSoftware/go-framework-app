// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
package settings

import (
	"fmt"

	"github.com/atrico-go/container"
)

type Settings interface {
	// Cmd line arguments
	GetArgument(name string) (value string, ok bool)
	MustGetArgument(name string) (value string)
	// Configuration is only read once (at startup)
	SingleReadConfiguration() bool
	// Dry run flag used
	IncludeDryRun() bool
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

type SetArgs interface {
	SetArgs(map[string]string)
}

// Register the settings
func RegisterSettings(c container.Container) {
	settings := theSettings{make(map[string]string)}
	c.Singleton(func() Settings { return &settings })
	c.Singleton(func() SetArgs { return &settings })
}

// Force all settings to be recalculated on next request
func ResetCaches() {
	singleReadConfigurationSettingCache.Reset()
	includeDryRunSettingCache.Reset()
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
type theSettings struct {
	args map[string]string
}

func (s *theSettings) SetArgs(args map[string]string) {
	s.args = args
}

func (s theSettings) GetArgument(name string) (value string, ok bool) {
	value, ok = s.args[name]
	return value, ok
}

func (s theSettings) MustGetArgument(name string) string {
	if value, ok := s.GetArgument("game"); ok {
		return value
	}
	panic(fmt.Sprintf("'%s' argument not found", name))
}
