{"Type":"Mixed"}
// {{.Comment}}
package unit_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/atrico-go/container"
	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
	"github.com/atrico-go/testing/random"
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/cobra"
	"{{.RepositoryPath}}/cmd"
	"{{.RepositoryPath}}/pkg"
	"{{.RepositoryPath}}/settings"
)

// SECTION-START: Options
{{ if gt (len .UserSettings) 0 }}
var rg = random.NewValueGenerator()
{{ range .UserSettings}}{{if gt (len .OptionTestCases) 0 }}
var Option{{.NameCode}} = OptionSet{
{{- range .OptionTestCases}}
	{{ . }},
{{- end}}
}
{{- end}}

{{- end}}
{{- end}}

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
{{- range .Commands}}{{- if not .NoImplementation}}
{{- $cmdName := .Name}}{{- $useName := commaList (quoted .SplitPath)}}{{- $args := concat .Args .OptionalArgs}}
{{- range $.UserSettings}}{{$settingName := .NameCode}}
{{- if and (.AppliesToCmdOrRoot $cmdName) (or (ne .Cmdline "") (ne .CmdlineShortcut ""))}}
{{- range .OptionTestCaseNames}}
	{Command: []string{ {{- $useName -}} }, Args: []string{ {{- commaList (quoted $args) -}} }, Options: []Option{Option{{$settingName}}["{{.}}"]}},
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
}

// SECTION-END

func addUserTests(tests []CmdlineTestCase) []CmdlineTestCase{
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
{{- range .UserSettings}}
{{- if or (ne .Cmdline "") (ne .CmdlineShortcut "")}}
	Assert(t).That(results["{{.NameCode}}"], is.DeepEqualTo(expected.{{.NameCode}}()), "{{.NameCode}}")
{{- end}}
{{- end}}
}

func resetCommand() *cobra.Command {
	// Re-register singletons
	c := container.NewContainer()
	settings.RegisterSettings(c)
	cmd.RegisterCmd(c)
	registerMockApiFactories(c)
	// Reset settings
	viperEx.Reset()
{{- if .SingleReadConfiguration}}
	settings.ResetCaches()
{{- end}}
	// Return root cmd
	var cmdFactory cmd.CommandFactory
	c.Make(&cmdFactory)
	return cmdFactory.Create()
}

// SECTION-END

