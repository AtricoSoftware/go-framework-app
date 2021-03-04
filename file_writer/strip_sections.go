package file_writer

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

type FilePart struct {
	Section  string
	Contents string
}

func StripSections(contents []byte) []FilePart {
	currentState := startState()
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	for scanner.Scan() {
		currentState = currentState.AddLine(scanner.Text() + "\n")
	}
	return currentState.GetParts()
}

var CommentSection = "-comment"
var ImportsSection = "-imports"

// ----------------------------------------------------------------------------------------------------------------------------
// State
// ----------------------------------------------------------------------------------------------------------------------------

type state interface {
	AddLine(line string) state
	GetParts() []FilePart
}

func startState() state {
	return noPart(make([]FilePart, 0))
}

type partDetails struct {
	parts   []FilePart
	current FilePart
}

// ----------------------------------------------------------------------------------------------------------------------------
// No Part
// ----------------------------------------------------------------------------------------------------------------------------

type noPart []FilePart

func (s noPart) AddLine(line string) state {
	stype, name := getLineType(line)
	newPart := FilePart{
		name,
		line,
	}
	switch stype {
	case normal, endImports:
		return notInSection{
			s,
			newPart,
		}
	case comment:
		return commentSection{
			s,
			newPart,
		}
	case startImports:
		return inImports{
			s,
			"",
		}
	case startSection:
		return inSection{
			s,
			newPart,
		}
	case endSection:
		panic("Section end before start")
	}
	panic("invalid section type")
}

func (s noPart) GetParts() []FilePart {
	return s
}
// ----------------------------------------------------------------------------------------------------------------------------
// In Section
// ----------------------------------------------------------------------------------------------------------------------------

type inSection partDetails

func (s inSection) AddLine(line string) state {
	stype, _ := getLineType(line)
	newPart := FilePart{
		s.current.Section,
		s.current.Contents + line,
	}
	switch stype {
	case normal, endImports:
		return inSection{
			s.parts,
			newPart,
		}
	case comment:
		panic("Comment within section")
	case startImports:
		return inImports{
			append(s.parts, s.current),
			s.current.Section,
		}
	case startSection:
		panic("Section start before end")
	case endSection:
		return noPart(append(s.parts, newPart))
	}
	panic("invalid section type")
}

func (s inSection) GetParts() []FilePart {
	panic("Missing Section end")
}
// ----------------------------------------------------------------------------------------------------------------------------
// Not in a section
// ----------------------------------------------------------------------------------------------------------------------------

type notInSection partDetails

func (s notInSection) AddLine(line string) state {
	stype, name := getLineType(line)
	newPart := FilePart{
		name,
		line,
	}
	switch stype {
	case normal, endImports:
		return notInSection{
			s.parts,
			FilePart{
				name,
				s.current.Contents + line,
			},
		}
	case comment:
		return commentSection{
			append(s.parts, s.current),
			newPart,
		}
	case startImports:
		return inImports{
			s.parts,
			"",
		}
	case startSection:
		return inSection{
			append(s.parts, s.current),
			newPart,
		}
	case endSection:
		panic("Section end before start")
	}
	panic("invalid section type")
}

func (s notInSection) GetParts() []FilePart {
	return append(s.parts, s.current)
}
// ----------------------------------------------------------------------------------------------------------------------------
// In a comment
// ----------------------------------------------------------------------------------------------------------------------------

type commentSection partDetails

func (s commentSection) AddLine(line string) state {
	stype, name := getLineType(line)
	newPart := FilePart{
		name,
		line,
	}
	switch stype {
	case normal, endImports:
		return notInSection{
			append(s.parts, s.current),
			newPart,
		}
	case comment:
		return commentSection{
			s.parts,
			FilePart{
				name,
				s.current.Contents + line,
			},
		}
	case startImports:
		return inImports{
			s.parts,
			"",
		}
	case startSection:
		return inSection{
			append(s.parts, s.current),
			newPart,
		}
	case endSection:
		panic("Section end before start")
	}
	panic("invalid section type")
}

func (s commentSection) GetParts() []FilePart {
	return append(s.parts, s.current)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Imports block
// ----------------------------------------------------------------------------------------------------------------------------

type inImports struct {
	parts   []FilePart
	parent string
}

func (s inImports) AddLine(line string) state {
	stype, _ := getLineType(line)
	switch stype {
	case normal, startImports:
		return s
	case  comment:
		panic("Comment within imports")
	case startSection, endSection:
		panic("Section within imports")
	case endImports:
		// Return to parent
		parts := append(s.parts, FilePart{
			Section:  ImportsSection,
			Contents: "",
		})
		if s.parent == "" {
			return noPart(parts)
		} else {
			return inSection{
				parts:   parts,
				current: FilePart{
					Section:  s.parent+"_extra",
					Contents: "",
				},
			}
		}
	}
	panic("invalid section type")
}

func (s inImports) GetParts() []FilePart {
	return s.parts
}

type lineType = int

const (
	normal       lineType = iota
	comment      lineType = iota
	startImports lineType = iota
	endImports   lineType = iota
	startSection lineType = iota
	endSection   lineType = iota
)

var sectionRegex = regexp.MustCompile(`SECTION-(START|END)(: (\w+))*`)

func getLineType(line string) (stype lineType, name string) {
	if FileCommentRegexp.MatchString(line) {
		return comment, CommentSection
	}
	if strings.HasPrefix(strings.TrimSpace(line), "import") {
		return startImports, ImportsSection
	}
	if strings.HasPrefix(strings.TrimSpace(line), ")") {
		return endImports, ""
	}
	if match := sectionRegex.FindStringSubmatch(line); match != nil {
		if match[1] == "START" {
			return startSection, match[3]
		} else {
			return endSection, ""
		}
	}
	return normal, ""
}
