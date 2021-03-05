package file_writer

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type ImportItem struct {
	Url   string
	Alias string
}

func (i ImportItem) String() string {
	str := strings.Builder{}
	if i.Alias != "" {
		str.WriteString(fmt.Sprintf("%s ", i.Alias))
	}
	str.WriteString(fmt.Sprintf(`"%s"`, i.Url))
	return str.String()
}

func AddImports(contents []byte, imports []ImportItem) []ImportItem {
	importsSet := make(map[ImportItem]interface{}, len(imports))
	for _, item := range imports {
		importsSet[item] = nil
	}
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	inImports := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if inImports {
			// End of imports?
			if strings.HasPrefix(line, ")") {
				break
			}
			// Read next import
			if line != "" {
				item := parseImportItem(line)
				if _, ok := importsSet[item]; !ok {
					importsSet[item] = nil
					imports = append(imports, item)
				}
			}
		} else {
			// Look for start of imports
			if strings.HasPrefix(line, "import") {
				if strings.Contains(line, "(") {
					inImports = true
				} else {
					imports = append(imports, parseImportItem(strings.TrimSpace(line[6:])))
				}
			}
		}
	}
	// Imports will be sorted later (by goimports)
	return imports
}

func FormatImports(imports []ImportItem) string {
	str := strings.Builder{}
	str.WriteString("import (\n")
	for _,item := range imports {
		str.WriteString(fmt.Sprintf("\t%s\n", item.String()))
	}
	str.WriteString(")\n")
	return str.String()
}

func parseImportItem(line string) (item ImportItem) {
	parts := strings.Split(line, " ")
	lenParts := len(parts)
	url := strings.Trim(parts[lenParts-1], `"`)
	alias := ""
	if lenParts > 1 {
		alias = parts[0]
	}
	return ImportItem{Url: url, Alias: alias}
}
