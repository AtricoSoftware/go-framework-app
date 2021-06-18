package api

import (
	"fmt"
	"io"
)

var VerboseFlag bool

func VerbosePrintln(a ...interface{}) (n int, err error) {
	if VerboseFlag {
		return fmt.Println(a...)
	} else {
		return 0, nil
	}
}

func VerbosePrintf(format string, a ...interface{}) (n int, err error) {
	if VerboseFlag {
		return fmt.Printf(format, a...)
	} else {
		return 0, nil
	}
}

func VerboseFprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if VerboseFlag {
		return fmt.Println(a...)
	} else {
		return 0, nil
	}
}

func VerboseFprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if VerboseFlag {
		return fmt.Printf(format, a...)
	} else {
		return 0, nil
	}
}