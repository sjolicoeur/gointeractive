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
	//


	screen.Display("Loading modules:", true, "tmp")
	screen.Display("a", true, "tmp")
	screen.Display("b", true, "tmp")
	time.Sleep(250 * time.Millisecond)
	screen.Display("c", true, "tmp")
	time.Sleep(450 * time.Millisecond)

	screen.Display("d", true, "tmp")
	time.Sleep(650 * time.Millisecond)
	screen.Display("f", true, "tmp")
	time.Sleep(1050 * time.Millisecond)

	screen.Display("z", true, "tmp")
	time.Sleep(2050 * time.Millisecond)
	screen.CarvePrint("Done Loading modules!")
	time.Sleep(2050 * time.Millisecond)
	screen.ClearNamedLayers("tmp")
	screen.Render()
	screen.CarvePrint("Goodbye!")





}
