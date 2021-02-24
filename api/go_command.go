package api

import (
	"os"
	"os/exec"
)

func GoCommand(targetDirectory string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = targetDirectory
	return cmd.Run()

}
