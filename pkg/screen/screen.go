package screen

import (
"fmt"

)

type Line struct {
content string
keep bool
}


func (line *Line) isBlank() bool {
return line.content == ""

} 

type Screen struct {
lines []Line
keepWhitespace bool
isATty bool // if true filer escape codes ?

}

func (s *Screen) Display( content string, preserveContent bool) error {
// break content with \n into lines
// call render?
}


func (s *Screen) ShowPrint ( content string) error {
return s.Display(content, false)
}

//
func (s *Screen) CarvePrint ( content string) error {
// or call the func carve
return s.Display(content, true)
}

func (s *Screen) Render() error {
// loop over lines to erase lines with keep set to false
// clear the screen
// rewrite the lines
}

func (s *Screen) CleanLines() {
// remove all lines set to keep == false
// mke this private?
}

func (s *Screen) RemoveBlankLines() error {
// loop over lines and remove lines that are blank
}

func (s *Screen) Clear() error {
// clear the actual screen so it can be repinted
}


// ability to insert at line X


/*

#!/usr/bin/env python

from __future__ import print_function

for x in range(0, 90):
    print(x);print(" - %s cows" % x )
    time.sleep(0.1)
    print("\033[A\033[2K\033[A\033[2K", end="")

*/
// may have issues with lines that wrap...? 
