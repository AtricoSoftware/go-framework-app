// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
// SECTION-START: Framework
package api

import (
	"fmt"
	"sort"
	"strings"

	"github.com/atrico-go/container"
	"github.com/atrico-go/core"
	"github.com/atrico-go/display"

	"github.com/AtricoSoftware/go-framework-app/settings"
)

type ListSkeletonsApi Runnable
type ListSkeletonsApiFactory Factory

type listSkeletonsApiFactory struct {
	container.Container
}

func (f listSkeletonsApiFactory) Create(args []string) Runnable {
	RegisterApiListSkeletons(f.Container)
	var theApi ListSkeletonsApi
	f.Container.Make(&theApi)
	return theApi
}

// SECTION-END

func RegisterApiListSkeletons(c container.Container) {
	c.Singleton(func(config settings.Settings, verboseService settings.VerboseService) ListSkeletonsApi {
		return listSkeletonsApi{config, verboseService}
	})
}

type listSkeletonsApi struct {
	settings.Settings
	settings.VerboseService
}

// List the available settings skeletons
func (svc listSkeletonsApi) Run() error {
	settings.ReadSkeletonFiles(svc.SkeletonFiles())
	table := display.NewTableBuilder().WithVerticalSeparator(' ')
	cells := []interface{}{
		"SKELETON",
		"CONFIG ID",
		"CMDLINE",
		"TYPE",
		"ENV VAR",
		"DEFAULT VAL",
		"DESCRIPTION",
	}
	table.AppendRow(cells...)
	for i, str := range cells {
		cells[i] = strings.Repeat("-", len(str.(string)))
	}
	table.AppendRow(cells...)
	names := make([]string, 0, len(settings.SkeletonCloset))
	for name := range settings.SkeletonCloset {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })
	for _, name := range names {
		cmdLine := strings.Builder{}
		s := settings.SkeletonCloset[name]
		separator := " "
		if s.CmdlineShortcut != "" {
			cmdLine.WriteString("-")
			cmdLine.WriteString(s.CmdlineShortcut)
			separator = ","
		} else {
			cmdLine.WriteString("  ")
		}
		if s.Cmdline != "" {
			cmdLine.WriteString(fmt.Sprintf("%s --%s", separator, s.Cmdline))
		}
		table.AppendRow(name, s.Id, cmdLine.String(), s.Type, s.EnvVar, s.DefaultVal, s.Description)
	}
	core.DisplayMultiline(table.Build().Render())
	return nil
}
