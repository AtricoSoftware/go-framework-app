package file_writer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (f fileWriter) RemoveObsoleteBackups() {
	// Remove backups with no changes
	for _, info := range *f.generatedFiles {
		if info.fullBackupPath() != "" {
			if filesEqual(info.fullOriginalPath(), info.fullBackupPath()) {
				// Files equal, remove backup
				os.Remove(info.fullBackupPath())
			}
		}
	}
}

func (f fileWriter) CleanupFiles() {
	for _, file := range *f.generatedFiles {
		if filepath.Ext(file.originalPath) == ".go" {
			goImports(f.config.RepositoryPath(), file.fullOriginalPath())
		}
	}
}

func goImports(local string, file string) error {
	cmd := exec.Command("goimports", "-w", "--local", local, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func fileExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return err == nil
}

// True if files are equal
// False if error
func filesEqual(path1, path2 string) bool {
	if file1, err := os.Open(path1); err == nil {
		defer file1.Close()
		if file2, err := os.Open(path2); err == nil {
			defer file2.Close()
			scanner1 := bufio.NewScanner(file1)
			scanner2 := bufio.NewScanner(file2)
			// Read each line
			for scanner1.Scan() && scanner2.Scan() {
				// Skip comment if present
				if FileCommentRegexp.MatchString(scanner1.Text()) && FileCommentRegexp.MatchString(scanner2.Text()) {
					continue
				}
				// Compare lines
				if scanner1.Text() != scanner2.Text() {
					return false
				}
			}
			return true
		}
	}
	return false
}

// Sprintf but extra/missing warnings stripped
func safeSprintf(format string, args ...interface{}) string {
	return strings.Split(fmt.Sprintf(format, args...), "%!")[0]
}
