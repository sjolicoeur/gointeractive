package main

import (
	//"fmt"
	"github.com/sjolicoeur/gointeractive/pkg/screen"
	"time"
)

func main() {
	//fmt.Print("\n")
	screen := screen.NewScreen()
	screen.CarvePrint("Welcome let's run this!")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("test")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("test.")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("test...")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("test......")
	time.Sleep(250 * time.Millisecond)
	screen.CarvePrint("test... done!")
	time.Sleep(250 * time.Millisecond)
	screen.CarvePrint("Done")

}
