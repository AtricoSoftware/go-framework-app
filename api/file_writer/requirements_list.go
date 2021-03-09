package file_writer

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type RequireItem struct {
	Url     string
	Version string
	Comment string
}

func (i RequireItem) String() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf(`%s %s`, i.Url, i.Version))
	if i.Comment != "" {
		str.WriteString(fmt.Sprintf(" // %s", i.Comment))
	}
	return str.String()
}

func AddRequirements(contents []byte, requirements []RequireItem) []RequireItem {
	requiresSet := make(map[string]RequireItem, len(requirements))
	for _, item := range requirements {
		requiresSet[item.Url] = item
	}
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	inRequires := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if inRequires {
			// End of requirements?
			if strings.HasPrefix(line, ")") {
				break
			}
			// Read next import
			if line != "" {
				item := parseRequireItem(line)
				if existing, ok := requiresSet[item.Url]; !ok || higherVersion(item.Version, existing.Version) {
					// New entry
					requiresSet[item.Url] = item
				}
			}
		} else {
			// Look for start of requirements
			if strings.HasPrefix(line, "require") {
				inRequires = true
			}
		}
	}
	newRequirements := make([]RequireItem, 0, len(requiresSet))
	for _, req := range requiresSet {
		newRequirements = append(newRequirements, req)
	}
	sort.Slice(newRequirements, func(i, j int) bool { return newRequirements[i].Url < newRequirements[j].Url })
	return newRequirements
}

func FormatRequirements(requirements []RequireItem) string {
	str := strings.Builder{}
	str.WriteString("require (\n")
	for _, item := range requirements {
		str.WriteString(fmt.Sprintf("\t%s\n", item.String()))
	}
	str.WriteString(")\n")
	return str.String()
}

var requireRegex = regexp.MustCompile(`(?P<url>\S+)\s+(?P<version>\S+)(?:$|\s*//\s*(?P<comment>.*))?`)

func parseRequireItem(line string) (item RequireItem) {
	values, _ := getNamedCaptureGroups(requireRegex, line)
	return RequireItem{Url: values["url"], Version: values["version"], Comment: values["comment"]}
}

var versionRegex = regexp.MustCompile(`v(?P<maj>\d+).(?P<min>\d+).(?P<patch>\d+)(?:-(?P<label>\S+))?`)

func higherVersion(lhs, rhs string) bool {
	lValues, _ := getNamedCaptureGroups(versionRegex, lhs)
	rValues, _ := getNamedCaptureGroups(versionRegex, rhs)
	var l uint64 = 0
	var r uint64 = 0
	for _, val := range []string{"maj", "min", "patch"} {
		l <<= 16
		l += uint64(AtoiOrDefault(lValues[val], 0))
		r <<= 16
		r += uint64(AtoiOrDefault(rValues[val], 0xffff))
	}
	l <<= 16
	if lValues["label"] != "" {
		l++
	}
	r <<= 16
	if rValues["label"] != "" {
		r++
	}
	return l > r
}

func AtoiOrDefault(str string, defaultVal int) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return defaultVal
}

func getNamedCaptureGroups(re *regexp.Regexp, text string) (values map[string]string, err error) {
	names := re.SubexpNames()[1:]
	values = make(map[string]string, len(names))
	if match := re.FindStringSubmatch(text); match != nil {
		for _, name := range names {
			values[name] = match[re.SubexpIndex(name)]
		}

	} else {
		for _, name := range names {
			values[name] = ""
		}
		err = errors.New("no match")
	}
	return values, err
}
