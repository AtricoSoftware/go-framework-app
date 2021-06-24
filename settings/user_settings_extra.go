package settings

func generateExtraSettings() (extra []UserSetting) {
	// Config file, Verbose
	return append(extra, configSetting, verboseSetting)
}

var configSetting = UserSetting{
	Name:                 "ConfigFile",
	Id:                   "ConfigFile",
	Description:          "Alternate config file",
	Type:                 "string",
	Cmdline:              "config-file",
	AppliesTo:            []string{"root"},
	optionTestCaseValues: emptyTcValues(),
}
var verboseSetting = UserSetting{
	Name:                 "Verbose",
	Id:                   "Verbose",
	Description:          "Generate more detailed output",
	Type:                 "bool",
	Cmdline:              "verbose",
	CmdlineShortcut:      "v",
	EnvVar:               "VERBOSE",
	AppliesTo:            []string{"root"},
	optionTestCaseValues: emptyTcValues(),
}
