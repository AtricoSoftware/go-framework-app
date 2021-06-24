// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
package unit_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/AtricoSoftware/go-framework-app/cmd"
	"github.com/AtricoSoftware/go-framework-app/pkg"
	"github.com/AtricoSoftware/go-framework-app/settings"
	"github.com/atrico-go/container"
	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
	"github.com/atrico-go/testing/random"
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/cobra"
)

// SECTION-START: Options

var rg = random.NewValueGenerator()

var OptionTargetDirectory = OptionSet{
	"Default": NewSimpleOption("--directory", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.TargetDirectoryVar = value.(string) }),
	"Short":   NewSimpleOption("-d", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.TargetDirectoryVar = value.(string) }),
}
var OptionApplicationTitle = OptionSet{
	"Default": NewSimpleOption("--title", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationTitleVar = value.(string) }),
	"Short":   NewSimpleOption("-t", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationTitleVar = value.(string) }),
}
var OptionApplicationName = OptionSet{
	"Default": NewSimpleOption("--name", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationNameVar = value.(string) }),
	"Short":   NewSimpleOption("-n", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationNameVar = value.(string) }),
}
var OptionApplicationSummary = OptionSet{
	"Default": NewSimpleOption("--summary", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationSummaryVar = value.(string) }),
}
var OptionApplicationDescription = OptionSet{
	"Default": NewSimpleOption("--description", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationDescriptionVar = value.(string) }),
}
var OptionRepositoryPath = OptionSet{
	"Default": NewSimpleOption("--repository", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.RepositoryPathVar = value.(string) }),
	"Short":   NewSimpleOption("-r", func() interface{} { var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.RepositoryPathVar = value.(string) }),
}
var OptionSkeletonFiles = OptionSet{
	"Default": NewSliceOption("--skeleton", func() interface{} { var value []string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { core.ConvertSlice(value, &s.SkeletonFilesVar) }),
	"Short":   NewSliceOption("-s", func() interface{} { var value []string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { core.ConvertSlice(value, &s.SkeletonFilesVar) }),
}
var OptionVerbose = OptionSet{
	"Default":     NewBooleanOption("--verbose", func(s *MockSettings) { s.VerboseVar = true }),
	"=True":       NewBooleanOption("--verbose=true", func(s *MockSettings) { s.VerboseVar = true }),
	"=False":      NewBooleanOption("--verbose=false", func(s *MockSettings) { s.VerboseVar = false }),
	"Short":       NewBooleanOption("-v", func(s *MockSettings) { s.VerboseVar = true }),
	"Short=True":  NewBooleanOption("-v=true", func(s *MockSettings) { s.VerboseVar = true }),
	"Short=False": NewBooleanOption("-v=false", func(s *MockSettings) { s.VerboseVar = false }),
}

// SECTION-END

// SECTION-START: TestCases
// ----------------------------------------------------------------------------------------------------------------------------
// Test cases
// ----------------------------------------------------------------------------------------------------------------------------
type CmdlineTestCase struct {
	Command []string
	Args    []string
	Options []Option
}

var CmdlineTestCases = []CmdlineTestCase{
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionTargetDirectory["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionTargetDirectory["Short"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionApplicationTitle["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionApplicationTitle["Short"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionApplicationName["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionApplicationName["Short"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionApplicationSummary["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionApplicationDescription["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionRepositoryPath["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionRepositoryPath["Short"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionSkeletonFiles["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionSkeletonFiles["Short"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionVerbose["Default"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionVerbose["=True"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionVerbose["=False"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionVerbose["Short"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionVerbose["Short=True"]}},
	{Command: []string{"generate"}, Args: []string{}, Options: []Option{OptionVerbose["Short=False"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionSkeletonFiles["Default"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionSkeletonFiles["Short"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionVerbose["Default"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionVerbose["=True"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionVerbose["=False"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionVerbose["Short"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionVerbose["Short=True"]}},
	{Command: []string{"list", "skeletons"}, Args: []string{}, Options: []Option{OptionVerbose["Short=False"]}},
}

// SECTION-END

func addUserTests(tests []CmdlineTestCase) []CmdlineTestCase {
	// Append extra tests here
	return tests
}

// SECTION-START: Test
// ----------------------------------------------------------------------------------------------------------------------------
// Test
// ----------------------------------------------------------------------------------------------------------------------------

func Test_CommandLine(t *testing.T) {
	for _, testCase := range addUserTests(CmdlineTestCases) {
		// Build command line and expectations
		cmdline := strings.Builder{}
		cmdline.WriteString(strings.Join(testCase.Command, " "))
		expected := NewMockSettings(testCase.Command, testCase.Args)
		for _, arg := range testCase.Args {
			cmdline.WriteString(" ")
			cmdline.WriteString(arg)
		}
		for _, opt := range testCase.Options {
			opt.Set()
			cmdline.WriteString(" ")
			cmdline.WriteString(opt.Cmdline())
			opt.ModifySettings(&expected)
		}
		t.Run(cmdline.String(), func(t *testing.T) { testCommandLineImpl(t, cmdline.String(), expected) })
	}
}

func testCommandLineImpl(t *testing.T, cmdline string, expected MockSettings) {
	// Arrange
	args := strings.Split(cmdline, " ")
	os.Args = make([]string, 1+len(args))
	os.Args[0] = pkg.Name
	copy(os.Args[1:], args)
	cmd := resetCommand()

	// Act
	err := cmd.Execute()

	// Assert
	Assert(t).That(err, is.EqualTo(nil), "Error")
	Assert(t).That(results["TheCommand"], is.DeepEqualTo(expected.TheCommand), "Command")
	Assert(t).That(results["Args"], is.DeepEqualTo(expected.TheArgs), "Args")
	Assert(t).That(results["TargetDirectory"], is.DeepEqualTo(expected.TargetDirectory()), "TargetDirectory")
	Assert(t).That(results["ApplicationTitle"], is.DeepEqualTo(expected.ApplicationTitle()), "ApplicationTitle")
	Assert(t).That(results["ApplicationName"], is.DeepEqualTo(expected.ApplicationName()), "ApplicationName")
	Assert(t).That(results["ApplicationSummary"], is.DeepEqualTo(expected.ApplicationSummary()), "ApplicationSummary")
	Assert(t).That(results["ApplicationDescription"], is.DeepEqualTo(expected.ApplicationDescription()), "ApplicationDescription")
	Assert(t).That(results["RepositoryPath"], is.DeepEqualTo(expected.RepositoryPath()), "RepositoryPath")
	Assert(t).That(results["SkeletonFiles"], is.DeepEqualTo(expected.SkeletonFiles()), "SkeletonFiles")
	Assert(t).That(results["ConfigFile"], is.DeepEqualTo(expected.ConfigFile()), "ConfigFile")
	Assert(t).That(results["Verbose"], is.DeepEqualTo(expected.Verbose()), "Verbose")
}

func resetCommand() *cobra.Command {
	// Re-register singletons
	c := container.NewContainer()
	settings.RegisterSettings(c)
	cmd.RegisterCmd(c)
	registerMockApiFactories(c)
	// Reset settings
	viperEx.Reset()
	settings.ResetCaches()
	// Return root cmd
	var cmdFactory cmd.CommandFactory
	c.Make(&cmdFactory)
	return cmdFactory.Create()
}

// SECTION-END
