package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sjolicoeur/gointeractive/pkg/screen"
)

var sentences = []string{
	"Shaving the cat",
	"Removing lint from the dryer",
	"baking burgers",
	"Reading the postits",
}

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
	time.Sleep(1050 * time.Millisecond)
	screen.CarvePrint("Done Loading modules!")
	time.Sleep(1050 * time.Millisecond)


	//fmt.Println("ba" + strings.Repeat("na", 2))
	r := rand.New(rand.NewSource(24))
	for i := 0; i <= 20; i++ {
		screen.ClearNamedLayers("buff")
		for _, sentence := range sentences {
			if i < 20 {
				dots := strings.Repeat(".", r.Intn(34))
				newLine := fmt.Sprintf("%-30s %-34s", sentence, dots)
				//newLine := sentence + strings.Repeat(".", r.Intn(34))
				screen.InsertLine(newLine, "buff")
			} else {
				newLine := fmt.Sprintf("> %-30s...%6s", sentence, "done!")
				//newLine := sentence + strings.Repeat(".", 10) + " done!"
				screen.InsertLine(newLine, "buff")
			}
		}
		screen.Render()
		time.Sleep(250 * time.Millisecond)

	}
	///
	screen.ClearNamedLayers("tmp")
	screen.Render()
	time.Sleep(1050 * time.Millisecond)
	screen.ClearNamedLayers("buff")
	screen.Render()
	screen.CarvePrint("Goodbye!")





}
