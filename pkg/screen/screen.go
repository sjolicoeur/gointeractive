package screen

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
)

// Screen is the global primitive that keeps track of
// lines on the screen and allow to operate o them to do
// things like animations or rendering
type Screen struct {
	lines          []Line
	keepWhitespace bool
	isATty         bool // if true filer escape codes ?
	au             aurora.Aurora
}

func NewScreen(showColors *bool) *Screen {
	lines := []Line{}
	au := aurora.NewAurora(*showColors)

	return &Screen{
		lines:          lines,
		keepWhitespace: false,
		isATty:         true,
		au:             au,
	}
}

// Clear clears the actual screen so it can be repinted
func (s *Screen) Clear() {
	for _, _ = range s.lines {
		fmt.Print("\033[A\033[2K")
	}
}

// Render renders all lines to the screen
func (s *Screen) Render() {
	for _, line := range s.lines {
		fmt.Println(line.content)
	}
}

// InsertLine adds a line to the screen without rendering or clearing the screen
func (s *Screen) InsertLine(content string, name string) {
	//s.Clear()
	line := NewLine(content, true, name)
	// append new lines to the lines
	s.lines = append(s.lines, *line)
}

// Display is the basic call to print a line to the screen.
// it allows to set the `keep` flag for the line and to name line
// successive calls to display will clear lines who have `keep` set to false
func (s *Screen) Display(content string, preserveContent bool, lineName string) {
	// break content with \n into lines
	tempLinesArr := strings.Split(content, "\n")
	// cleanup previous lines
	s.Clear()
	// remove lines we do not want to keep
	s.CleanLines()
	for _, tmpLine := range tempLinesArr {
		line := NewLine(tmpLine, preserveContent, lineName)
		s.lines = append(s.lines, *line)
	}
	s.Render()
}

// ShowPrint is a shortcut to an unnamed temporary line
func (s *Screen) ShowPrint(content string) {
	s.Display(content, false, "")
}

// CarvePrint is a shortcut to an unnamed permanent line
func (s *Screen) CarvePrint(content string) {
	s.Display(content, true, "")
}

func (s *Screen) CleanLines() {
	var tmpLines []Line
	for _, line := range s.lines {
		if line.keep == true {
			tmpLines = append(tmpLines, line)
		}
	}
	s.lines = tmpLines
}

func (s *Screen) RemoveBlankLines() {
	var tmpLines []Line
	for _, line := range s.lines {
		if line.isBlank() != true {
			tmpLines = append(tmpLines, line)
		}
	}
	s.lines = tmpLines
}

// ClearNamedLayers clears the layers with the specified name
func (s *Screen) ClearNamedLayers(layerName string) {
	s.Clear()
	var tmpLines []Line
	for _, line := range s.lines {
		if line.name != layerName {
			tmpLines = append(tmpLines, line)
		}
	}
	s.lines = tmpLines
}

// NumLines returns the number of lines that are known by screen
func (s *Screen) NumLines() int {
	return len(s.lines)
}

// Ok renders text in green
func (s *Screen) Ok(text string) string {
	return aurora.Sprintf(s.au.Green(text))
}

// Warning renders text in yellow
func (s *Screen) Warning(text string) string {
	return aurora.Sprintf(s.au.Brown(text))
}

// Critical renders text in red
func (s *Screen) Critical(text string) string {
	return aurora.Sprintf(s.au.Red(text))
}

// Bleu renders text in green
func (s *Screen) Bleu(text string) string {
	return aurora.Sprintf(s.au.Blue(text))
}

// Purple renders text in purple
func (s *Screen) Purple(text string) string {
	return aurora.Sprintf(s.au.Magenta(text))
}

// Teal renders text in teal color
func (s *Screen) Teal(text string) string {
	return aurora.Sprintf(s.au.Cyan(text))
}

// Normal renders text as normal is a noop
func (s *Screen) Normal(text string) string {
	return text
}

// Emphasis emboldens the text
func (s *Screen) Emphasis(text string) string {
	return aurora.Sprintf(s.au.Bold(text))
}

// Invert inverts the color of the text
// practical for when the background color is changed
func (s *Screen) Invert(text string) string {
	return aurora.Sprintf(s.au.Inverse(text))
}
