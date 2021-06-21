package settings

func generateExtraSettings() (extra []UserSetting) {
	// Verbose
	extra = append(extra, verboseSetting)

	return extra
}

var verboseSetting = UserSetting{
	Name:                 "Verbose",
	Id:                   "Verbose",
	Description:          "Generate more detailed output",
	Type:                 "bool",
	Cmdline:              "verbose",
	EnvVar:               "VERBOSE",
	DefaultVal:           "false",
	AppliesTo:            []string{"root"},
	optionTestCaseValues: emptyTcValues(),
}
