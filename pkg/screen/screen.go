package screen

import (
	"fmt"
)



type Screen struct {
	lines          []Line
	keepWhitespace bool
	isATty         bool // if true filer escape codes ?
}

func NewScreen() *Screen {
	lines := []Line{}

	return &Screen{
		lines: lines,
		keepWhitespace: false,
		isATty: true,
	}
}

func (s *Screen) Display(content string, preserveContent bool, lineName string) error {
	// break content with \n into lines
	// cleanup previous lines
	s.Clear()
	// remove lines we do not want to keep
	s.CleanLines()
	// call render?
	line := NewLine(content, preserveContent, lineName)
	// append new lines to the lines
	s.lines = append(s.lines, *line)
	// print lines
	s.Render()
	//
	return nil
}

func (s *Screen) ShowPrint(content string) error {
	return s.Display(content, false, "")
}

//
func (s *Screen) CarvePrint(content string) error {
	// or call the func carve
	return s.Display(content, true, "")
}

func (s *Screen) Render() error {
	// loop over lines to erase lines with keep set to false
	// clear the screen
	// rewrite the lines
	for _, line := range s.lines {
		fmt.Println(line.content)
	}
	return nil
}

func (s *Screen) CleanLines() error {
	// remove all lines set to keep == false
	// mke this private?
	var tmpLines []Line
	for _, line := range s.lines {
		if line.keep == true {
			tmpLines = append(tmpLines, line)
		}
	}
	s.lines = tmpLines
	return nil
}

func (s *Screen) RemoveBlankLines() error {
	// loop over lines and remove lines that are blank
	return nil
}

func (s *Screen) Clear() error {
	// clear the actual screen so it can be repinted
	//numLines := len(s.lines)
	for _, _ = range s.lines {
		fmt.Print("\033[A\033[2K")
	}
	return nil
}

func (s *Screen) ClearNamedLayers(layerName string) error {
	// clear the layers based on a name
	s.Clear()
	var tmpLines []Line
	for _, line := range s.lines {
		if line.name != layerName {
			tmpLines = append(tmpLines, line)
		}
	}
	s.lines = tmpLines
	return nil
}

// ability to insert at line X

/*

#!/usr/bin/env python

from __future__ import print_function
import time


for x in range(0, 90):
    print(x);print(" - %s cows" % x )
    time.sleep(0.1)
    print("\033[A\033[2K\033[A\033[2K", end="")

*/
// may have issues with lines that wrap...?
