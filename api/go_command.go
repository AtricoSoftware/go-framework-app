package api

import (
	"os"
	"os/exec"
)

func GoCommand(targetDirectory string, args ...string) error {
	return ExecuteCommand(targetDirectory, "go", args...)
}

func ExecuteCommand(targetDirectory string, exe string, args ...string) error {
	cmd := exec.Command(exe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = targetDirectory
	return cmd.Run()
}
