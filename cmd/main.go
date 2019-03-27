package main

import (
	//"fmt"
	"github.com/sjolicoeur/gointeractive/pkg/screen"
	"time"
)

func main() {
	//fmt.Print("\n")
	screen := screen.NewScreen()
	screen.Display("Welcome let's run this!", true, "")
	time.Sleep(250 * time.Millisecond)
	screen.Display("test", false, "")
	time.Sleep(250 * time.Millisecond)
	screen.Display("test.", false, "")
	time.Sleep(250 * time.Millisecond)
	screen.Display("test...", false, "")
	time.Sleep(250 * time.Millisecond)
	screen.Display("test......", false, "")
	time.Sleep(250 * time.Millisecond)
	screen.Display("test... done!", true, "")
	time.Sleep(250 * time.Millisecond)
	screen.Display("Done", true, "")

}
