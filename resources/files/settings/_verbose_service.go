package settings

import(
	"fmt"
	"io"
	"os"

	"github.com/atrico-go/container"
	"github.com/spf13/viper"
)

type VerboseService interface {
	VerbosePrintln(a ...interface{}) (n int, err error)
	VerbosePrintf(format string, a ...interface{}) (n int, err error)
	VerboseFprintln(w io.Writer, a ...interface{}) (n int, err error)
	VerboseFprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
}

func RegisterVerboseService(c container.Container) {
	c.Singleton(func(config Settings) VerboseService { return verboseService(config.Verbose()) })
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

// For internal use only
func GetVerboseService() VerboseService {
	return verboseService(viper.GetBool(verboseSettingName))
}

type verboseService bool

func (v verboseService) VerbosePrintln(a ...interface{}) (n int, err error) {
	return v.VerboseFprintln(os.Stdout, a...)
}

func (v verboseService) VerbosePrintf(format string, a ...interface{}) (n int, err error) {
	return v.VerboseFprintf(os.Stdout, format, a...)
}

func (v verboseService) VerboseFprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if v {
		return fmt.Println(a...)
	} else {
		return 0, nil
	}
}

func (v verboseService) VerboseFprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if v {
		return fmt.Printf(format, a...)
	} else {
		return 0, nil
	}
}

