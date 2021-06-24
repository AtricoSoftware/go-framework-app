// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
// SECTION-START: Framework
package api

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AtricoSoftware/go-framework-app/settings"
	"github.com/atrico-go/container"
	"github.com/atrico-go/core"
	"github.com/atrico-go/display"
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
	c.Singleton(func(config settings.Settings, verboseService settings.VerboseService) ListSkeletonsApi { return listSkeletonsApi{config, verboseService} })
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
		cmdLine := make([]string, 0, 2)
		s := settings.SkeletonCloset[name]
		if s.CmdlineShortcut != "" {
			cmdLine = append(cmdLine, fmt.Sprintf("-%s", s.CmdlineShortcut))
		}
		if s.Cmdline != "" {
			cmdLine = append(cmdLine, fmt.Sprintf("--%s", s.Cmdline))
		}
		table.AppendRow(name, s.Id, strings.Join(cmdLine, ", "), s.Type, s.EnvVar, s.DefaultVal, s.Description)
	}
	core.DisplayMultiline(table.Build().Render())
	return nil
}
