package viperEx

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var flags = make(map[string]*pflag.Flag)

// Used in testing to clear settings
func Reset() {
	viper.Reset()
	flags = make(map[string]*pflag.Flag)
}

func AddBoolSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddBoolSettingP(flagSet, name, "", description)
}

func AddBoolSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.BoolP(name, shorthand, false, description) })
}

func AddStringSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddStringSettingP(flagSet, name, "", description)
}

func AddStringSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.StringP(name, shorthand, "", description) })
}

func AddStringArraySetting(flagSet *pflag.FlagSet, name string, description string) {
	AddStringArraySettingP(flagSet, name, "", description)
}

func AddStringArraySettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.StringArrayP(name, shorthand, []string{}, description) })
}

func addSetting(flagSet *pflag.FlagSet, name string, createFlag func()) {
	if flag, ok := flags[name]; ok {
		// TODO [Improvement] - check type is the same
		// Add existing flag
		flagSet.AddFlag(flag)
	} else {
		// Create new flag
		createFlag()
		flag = flagSet.Lookup(name)
		// Bind to viper
		if err := viper.BindPFlag(name, flag); err != nil {
			log.Fatal("Unable to bind flag:", err)
		}
		// Store for next time
		flags[name] = flag
	}
}
