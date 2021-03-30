// Generated 2021-03-30 15:32:41 by go-framework development-version
package unit_tests

import (
	"os"
	"strings"
	"testing"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
	"github.com/AtricoSoftware/go-framework-app/cmd"
	"github.com/AtricoSoftware/go-framework-app/pkg"
	"github.com/AtricoSoftware/go-framework-app/settings"
	"github.com/atrico-go/container"
	"github.com/atrico-go/testing/random"
	"github.com/atrico-go/viperEx"
	"github.com/spf13/cobra"
)

// SECTION-START: Options

var rg = random.NewValueGenerator()

var OptionTargetDirectory = OptionSet {
	"Default": NewSimpleOption("--directory", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.TargetDirectoryVar = value.(string)}),
	"Short": NewSimpleOption("-d", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.TargetDirectoryVar = value.(string)}),
}
var OptionApplicationTitle = OptionSet {
	"Default": NewSimpleOption("--title", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationTitleVar = value.(string)}),
	"Short": NewSimpleOption("-t", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationTitleVar = value.(string)}),
}
var OptionApplicationName = OptionSet {
	"Default": NewSimpleOption("--name", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationNameVar = value.(string)}),
	"Short": NewSimpleOption("-n", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationNameVar = value.(string)}),
}
var OptionApplicationSummary = OptionSet {
	"Default": NewSimpleOption("--summary", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationSummaryVar = value.(string)}),
}
var OptionApplicationDescription = OptionSet {
	"Default": NewSimpleOption("--description", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.ApplicationDescriptionVar = value.(string)}),
}
var OptionRepositoryPath = OptionSet {
	"Default": NewSimpleOption("--repository", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.RepositoryPathVar = value.(string)}),
	"Short": NewSimpleOption("-r", func() interface{} {var value string; rg.Value(&value); return value }, func(s *MockSettings, value interface{}) { s.RepositoryPathVar = value.(string)}),
}
// SECTION-END
// SECTION-START: TestCases
// ----------------------------------------------------------------------------------------------------------------------------
// Test cases
// ----------------------------------------------------------------------------------------------------------------------------
type CmdlineTestCase struct {
	Command string
	Options []Option
}
var CmdlineTestCases = []CmdlineTestCase{
	{Command: "generate", Options: []Option {	OptionTargetDirectory["Default"] }},
	{Command: "generate", Options: []Option {	OptionTargetDirectory["Short"] }},
	{Command: "generate", Options: []Option {	OptionApplicationTitle["Default"] }},
	{Command: "generate", Options: []Option {	OptionApplicationTitle["Short"] }},
	{Command: "generate", Options: []Option {	OptionApplicationName["Default"] }},
	{Command: "generate", Options: []Option {	OptionApplicationName["Short"] }},
	{Command: "generate", Options: []Option {	OptionApplicationSummary["Default"] }},
	{Command: "generate", Options: []Option {	OptionApplicationDescription["Default"] }},
	{Command: "generate", Options: []Option {	OptionRepositoryPath["Default"] }},
	{Command: "generate", Options: []Option {	OptionRepositoryPath["Short"] }},
}

// SECTION-END

// SECTION-START: Test
// ----------------------------------------------------------------------------------------------------------------------------
// Test
// ----------------------------------------------------------------------------------------------------------------------------

func Test_CommandLine(t *testing.T) {
	for _,testCase := range addUserTests(CmdlineTestCases) {
		// Build command line and expectations
		cmdline := strings.Builder{}
		cmdline.WriteString(testCase.Command)
		expected := NewMockSettings(testCase.Command)
		for _, opt := range testCase.Options {
			opt.Set()
			cmdline.WriteString(" ")
			cmdline.WriteString(opt.Cmdline())
			opt.ModifySettings(&expected)
		}
		t.Run(cmdline.String(), func(t *testing.T) {testCommandLineImpl(t, cmdline.String(), expected)})
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
	Assert(t).That(results["TheCommand"], is.EqualTo(expected.TheCommand), "Command")
	Assert(t).That(results["TargetDirectory"], is.DeepEqualTo(expected.TargetDirectory()), "TargetDirectory")
	Assert(t).That(results["ApplicationTitle"], is.DeepEqualTo(expected.ApplicationTitle()), "ApplicationTitle")
	Assert(t).That(results["ApplicationName"], is.DeepEqualTo(expected.ApplicationName()), "ApplicationName")
	Assert(t).That(results["ApplicationSummary"], is.DeepEqualTo(expected.ApplicationSummary()), "ApplicationSummary")
	Assert(t).That(results["ApplicationDescription"], is.DeepEqualTo(expected.ApplicationDescription()), "ApplicationDescription")
	Assert(t).That(results["RepositoryPath"], is.DeepEqualTo(expected.RepositoryPath()), "RepositoryPath")
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
