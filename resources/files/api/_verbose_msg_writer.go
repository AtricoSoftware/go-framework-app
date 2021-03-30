package api

import (
	"fmt"
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
