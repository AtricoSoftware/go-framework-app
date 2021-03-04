package file_writer

import (
	"fmt"
	"regexp"
	"time"

	"github.com/AtricoSoftware/go-framework-app/pkg"
)

func FileComment(now time.Time) string {
	return fmt.Sprintf("Generated %s by %s %s", now.Format("2006-01-02 15:04:05"), pkg.Name, pkg.Version)
}

var FileCommentRegexp = regexp.MustCompile(fmt.Sprintf(`Generated \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} by %s `, pkg.Name))
