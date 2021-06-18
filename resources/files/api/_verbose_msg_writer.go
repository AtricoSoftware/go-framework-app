package api

import (
	"fmt"
	"io"
	"os"
)

var VerboseFlag bool

func VerbosePrintln(a ...interface{}) (n int, err error) {
	return VerboseFprintln(os.Stdout, a...)
}

func VerbosePrintf(format string, a ...interface{}) (n int, err error) {
	return VerboseFprintf(os.Stdout, format, a...)
}

func VerboseFprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if VerboseFlag {
		return fmt.Fprintln(w, a...)
	} else {
		return 0, nil
	}
}

func VerboseFprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if VerboseFlag {
		return fmt.Fprintf(w, format, a...)
	} else {
		return 0, nil
	}
}