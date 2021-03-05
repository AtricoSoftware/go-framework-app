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
var RequiresSection = "-requires"

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
	case normal, endMergeBlock:
		return notInSection{
			s,
			newPart,
		}
	case comment:
		return commentSection{
			s,
			newPart,
		}
	case startMergeBlock:
		return inMergeBlock{
		name,
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
	stype, name := getLineType(line)
	newPart := FilePart{
		s.current.Section,
		s.current.Contents + line,
	}
	switch stype {
	case normal, endMergeBlock:
		return inSection{
			s.parts,
			newPart,
		}
	case comment:
		panic("Comment within section")
	case startMergeBlock:
		return inMergeBlock{
		name,
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
	case normal, endMergeBlock:
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
	case startMergeBlock:
		return inMergeBlock{
		name,
			append(s.parts, s.current),
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
	case normal, endMergeBlock:
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
	case startMergeBlock:
		return inMergeBlock{
		name,
			append(s.parts, s.current),
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

type inMergeBlock struct {
	blockType string
	parts  []FilePart
	parent string
}

func (s inMergeBlock) AddLine(line string) state {
	stype, _ := getLineType(line)
	switch stype {
	case normal, startMergeBlock:
		return s
	case comment:
		panic("Comment within imports")
	case startSection, endSection:
		panic("Section within imports")
	case endMergeBlock:
		// Return to parent
		parts := append(s.parts, FilePart{
			Section:  s.blockType,
			Contents: "",
		})
		if s.parent == "" {
			return noPart(parts)
		} else {
			return inSection{
				parts: parts,
				current: FilePart{
					Section:  s.parent + "_extra",
					Contents: "",
				},
			}
		}
	}
	panic("invalid section type")
}

func (s inMergeBlock) GetParts() []FilePart {
	return s.parts
}

type lineType = int

const (
	normal          lineType = iota
	comment         lineType = iota
	startSection    lineType = iota
	endSection      lineType = iota
	startMergeBlock lineType = iota
	endMergeBlock   lineType = iota
)

var sectionRegex = regexp.MustCompile(`SECTION-(START|END)(: (\w+))*`)

func getLineType(line string) (stype lineType, name string) {
	if FileCommentRegexp.MatchString(line) {
		return comment, CommentSection
	}
	if strings.HasPrefix(strings.TrimSpace(line), "import") {
		return startMergeBlock, ImportsSection
	}
	if strings.HasPrefix(strings.TrimSpace(line), "require") {
		return startMergeBlock, RequiresSection
	}
	if strings.HasPrefix(strings.TrimSpace(line), ")") {
		return endMergeBlock, ""
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
