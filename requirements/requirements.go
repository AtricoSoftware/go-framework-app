package requirements

import (
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/common"
)

var requirements = []string{
	"github.com/atrico-go/container",
	"github.com/atrico-go/viperEx",
	"github.com/mitchellh/go-homedir",
	"github.com/spf13/cobra",
	"github.com/spf13/pflag",
	"github.com/spf13/viper",
}

func GetRequirements(targetDirectory string) {
	for _, req := range requirements {
		common.GoCommand(targetDirectory, "get", "-t", "-u", req)
	}
}
