package file_writer

import (
	"bufio"
	"os"
	"path/filepath"
	"time"
)

func CleanupBackups(generatedFiles []GeneratedFileInfo) {
	// Remove backups with no changes
	for _, info := range generatedFiles {
		if info.backupPath != "" {
			if filesEqual(info.originalPath, info.backupPath, info.comment) {
				// Files equal, remove backup
				os.Remove(info.backupPath)
			}
		}
	}
}

// Time of run (used for comments)
var runTime = time.Now()

func fileExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return err == nil
}

// True if files are equal
// False if error
func filesEqual(path1, path2 string, skipComment bool) bool {
	if file1, err := os.Open(path1); err == nil {
		defer file1.Close()
		if file2, err := os.Open(path2); err == nil {
			defer file2.Close()
			scanner1 := bufio.NewScanner(file1)
			scanner2 := bufio.NewScanner(file2)
			// Read each line
			for scanner1.Scan() && scanner2.Scan() {
				// Skip first line?
				if skipComment {
					skipComment = false
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

func getComment(filename string) string {
	if filename == ".gitignore" {
		return "# %s"
	}
	if filename == "go.mod" {
		return "// %s"
	}
	switch filepath.Ext(filename) {
	case ".go":
		return "// %s"
	case ".sh", ".yaml", ".yml":
		return "# %s"
	case ".md":
		return "[comment]: <> ( %s )"
	}
	return ""
}
