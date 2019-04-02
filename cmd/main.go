package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sjolicoeur/gointeractive/pkg/screen"
	//"github.com/sjolicoeur/gointeractive/pkg/formating"
)

type Loader struct {
	states      []string
	currentSate int
}

func (loader *Loader) Next() string {
	loader.currentSate = loader.currentSate + 1
	if loader.currentSate > (len(loader.states) - 1) {
		loader.currentSate = 0
	}
	return loader.states[loader.currentSate]
}

func NewLoader(steps []string) *Loader {
	return &Loader{
		states:      steps,
		currentSate: 0,
	}
}

var sentences = []string{
	"Shaving the cat",
	"Removing lint from the dryer",
	"baking burgers",
	"Reading the postits",
}

var colors = flag.Bool("nocolors", true, "enable or disable colors")


func main() {
	flag.Parse()
	fmt.Println("")

	spinner := NewLoader([]string{"|", "/", "-", "\\"})
	//loader := NewLoader([]string{"\u2809", "\u2812", "\u2824", "\u28C0", "\u2824", "\u2812",})
	loaderBraille := NewLoader([]string{"\u2801", "\u2809", "\u2819", "\u281B", "\u281F", "\u283F", "\u28BF", "\u28FF",})
	//loader := NewLoader([]string{"\u2800", "\u2836", "\u28FF", "\u2836", "\u2800",})
	loader := NewLoader([]string{
		"------",
		"o-----",
		"-o----",
		"--o---",
		"---o--",
		"----o-",
		"-----o",
		"------",
		"-----o",
		"----o-",
		"---o--",
		"--o---",
		"-o----",
		"o-----",
	})
	//fmt.Print("\n")
	screen := screen.NewScreen(colors)

	var colors = []func(string) string{
		screen.Ok,
		screen.Warning,
		screen.Critical,
		screen.Bleu,
		screen.Purple,
		screen.Teal,
	}

	var formats = []func(string) string{
		//screen.Invert,
		screen.Normal,
		screen.Emphasis,
	}

	//screen.ShowPrint("Welcome to this short demo!")
	//screen.InsertLine("Welcome to this short demo!", "narration")
	screen.Display("Welcome to this short demo!", true, "narration")
	//screen.Render()
	time.Sleep(250 * time.Millisecond)
	//screen.ShowPrint("Sit back and enjoy!")
	//screen.ClearNamedLayers("buff")
	//screen.InsertLine("Sit back and enjoy!", "narration")
	screen.Display("Sit back and enjoy!\nMaybe get popcorn?", true, "narration")

	//screen.Render()
	time.Sleep(1550 * time.Millisecond)
	//screen.CarvePrint()
	//screen.ClearNamedLayers("narration")
	//screen.CarvePrint("Let's start!")
	screen.Display("Let's start!", true, "narration")

	time.Sleep(250 * time.Millisecond)



	screen.ShowPrint("This library if meant to offer the primitives for something.")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("This library if meant to offer the primitives for something.")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("This library if meant to offer the primitives for something...")
	time.Sleep(250 * time.Millisecond)
	screen.ShowPrint("This library if meant to offer the primitives for something......")
	time.Sleep(250 * time.Millisecond)
	//screen.Clear()
	//screen.InsertLine(screen.Ok("This library if meant to offer the primitives for something... done!"), "narration")
	screen.Display(screen.Emphasis("This library if meant to offer the primitives for something... greater!"), true, "narration")


	//screen.Render()

	time.Sleep(250 * time.Millisecond)
	//screen.CarvePrint("Done")
	//screen.InsertLine("Done", "narration")
	screen.Display("With the help of ShowPrint() or CarvePrint().", true, "narration")
	screen.Display("You can Choose to have things:", true, "narration")
	screen.ShowPrint("- Shown for a brief moment")
	time.Sleep(750 * time.Millisecond)
	screen.CarvePrint("- Or Etched into the output")


	screen.Display("ShowPrint() is enough to perform animations on a single line. ", true, "narration")
	for i := 0; i <=  25; i++ {

		//screen.ClearNamedLayers("animation1")
		//screen.InsertLine("Like so:  "+loader.Next(), "animation1")
		//screen.Render()
		screen.ShowPrint("Like so:  "+spinner.Next())

		time.Sleep(75 * time.Millisecond)

	}
	//screen.ClearNamedLayers("animation1")
	//screen.Render()




	screen.Display("With the aid of the lower level Display(),", true, "narration")
	screen.Display("and it's ability to name layers.", true, "narration")
	screen.Display("We can do animations over multiple lines.", true, "narration")

	//screen.Render()
	//

	screen.Display("Animating top down, like so:", true, "tmp")
	screen.Display("* ", true, "tmp")
	screen.Display("**", true, "tmp")
	time.Sleep(250 * time.Millisecond)
	screen.Display("***", true, "tmp")
	time.Sleep(450 * time.Millisecond)

	screen.Display("****", true, "tmp")
	time.Sleep(650 * time.Millisecond)
	screen.Display("*****", true, "tmp")
	time.Sleep(550 * time.Millisecond)

	screen.Display("******", true, "tmp")
	time.Sleep(650 * time.Millisecond)
	//screen.CarvePrint("Done Loading modules!")
	//screen.Clear()
	screen.Display("******** Done!", true, "tmp")
	screen.Display("Naming the lines allows us to selectively erase content.", true, "narration")
	//screen.Render()
	time.Sleep(1050 * time.Millisecond)
	screen.ClearNamedLayers("tmp")
	screen.Render()
	screen.Display("Do this by calling ClearNamedLayers() and then Render().", true, "narration")

	screen.Display("Combining those with InsertLine allows us to go wild:", true, "narration")

	//fmt.Println("ba" + strings.Repeat("na", 2))
	r := rand.New(rand.NewSource(60))
	for i := 0; i <= 100; i++ {

		screen.ClearNamedLayers("buff")
		screen.InsertLine("loading  "+loaderBraille.Next(), "buff")
		for _, sentence := range sentences {
			//randColorFunc := colors[rand.Intn(len(colors))]
			randFormatFunc := formats[rand.Intn(len(formats))]
			if i < 90 {
				dots := strings.Repeat(".", r.Intn(34))
				newLine := randFormatFunc(fmt.Sprintf("%-30s %-51s", sentence, dots))
				//newLine := sentence + strings.Repeat(".", r.Intn(34))
				screen.InsertLine(newLine, "buff")
			} else {
				newLine := screen.Emphasis(fmt.Sprintf("> %-30s...%6s", sentence, "done!"))
				//newLine := sentence + strings.Repeat(".", 10) + " done!"
				screen.InsertLine(newLine, "buff")
			}
		}
		screen.Render()
		time.Sleep(50 * time.Millisecond)

	}

	screen.Display("And we also have color:", true, "narration")
	time.Sleep(1050 * time.Millisecond)
	screen.ClearNamedLayers("buff")
	screen.Render()

	//r := rand.New(rand.NewSource(24))
	for i := 0; i <= 150; i++ {

		screen.ClearNamedLayers("buff")
		screen.InsertLine("loading  "+loader.Next(), "buff")
		for _, sentence := range sentences {
			randColorFunc := colors[rand.Intn(len(colors))]
			randFormatFunc := formats[rand.Intn(len(formats))]
			if i < 140 {
				dots := strings.Repeat(".", r.Intn(34))
				newLine := randFormatFunc(randColorFunc(fmt.Sprintf("%-30s %-51s", sentence, dots)))
				//newLine := sentence + strings.Repeat(".", r.Intn(34))
				screen.InsertLine(newLine, "buff")
			} else {
				newLine :=  screen.Emphasis(screen.Ok(fmt.Sprintf("> %-30s...%6s", sentence, "done!")))
				//newLine := sentence + strings.Repeat(".", 10) + " done!"
				screen.InsertLine(newLine, "buff")
			}
		}
		screen.Render()
		time.Sleep(50 * time.Millisecond)

	}



	///
	screen.Display("It's time to end the demo.", true, "narration")
	time.Sleep(1050 * time.Millisecond)

	screen.ClearNamedLayers("buff")
	screen.Render()
	time.Sleep(1050 * time.Millisecond)
	screen.ClearNamedLayers("narration")
	screen.Render()
	screen.CarvePrint("Goodbye! And thank you!")

}
