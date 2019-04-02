package screen

import (
	"io"
	"strings"

	"os"
	"testing"
)

func TestDisplayToKeep(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
		numLines  int
	)
	showColor = false
	screen = NewScreen(&showColor)
	screen.InsertLine("a test", "do-not-delete")
	screen.Display("tester line", true, "delete")
	screen.CarvePrint(".")
	numLines = screen.NumLines()
	screen.Clear()

	if numLines != 3 {
		t.Errorf("Got: %d, wanted: %d.", numLines, 2)
	}
	screen.ShowPrint("another test")
	screen.Display("", false, "delete")
	screen.CleanLines()
	numLines = screen.NumLines()
	if numLines != 3 {
		t.Errorf("Got: %d, wanted: %d.", numLines, 2)
	}
	screen.Clear()
}

func TestDisplayTmpLine(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
	)

	oldStdout := os.Stdout
	r, writeFile, err := os.Pipe()
	if err != nil {
		t.Errorf("FAIL")
	}

	os.Stdout = writeFile
	//
	outC := make(chan string)
	go func() {
		var buf strings.Builder
		_, err := io.Copy(&buf, r)
		_ = r.Close()
		if err != nil {
			t.Errorf("FAIL")
			os.Exit(1)
		}
		outC <- buf.String()
	}()

	//
	showColor = true
	screen = NewScreen(&showColor)

	screen.ShowPrint("line 1")
	screen.Display("line 2", false, "")
	screen.CarvePrint("line 3")
	//screen.Render()

	defer func() {
		// Close pipe, restore stdout, get output.
		writeFile.Close()
		os.Stdout = oldStdout
		out := <-outC

		//var fail string
		err := recover()
		got := strings.TrimSpace(out)
		want := "line 1\n\x1b[A\x1b[2Kline 2\n\x1b[A\x1b[2Kline 3"

		if got != want {
			t.Errorf("Got: %#v, wanted: %#v.", got, want)
		}

		if err != nil {
			panic(err)
		}
	}()
}

func TestRender(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
	)

	oldStdout := os.Stdout
	r, writeFile, err := os.Pipe()
	if err != nil {
		t.Errorf("FAIL")
	}

	os.Stdout = writeFile
	//
	outC := make(chan string)
	go func() {
		var buf strings.Builder
		_, err := io.Copy(&buf, r)
		_ = r.Close()
		if err != nil {
			t.Errorf("FAIL")
			os.Exit(1)
		}
		outC <- buf.String()
	}()

	//
	showColor = true
	screen = NewScreen(&showColor)

	screen.InsertLine("line 1", "")
	screen.InsertLine("line 2", "")
	screen.Render()

	defer func() {
		// Close pipe, restore stdout, get output.
		writeFile.Close()
		os.Stdout = oldStdout
		out := <-outC

		//var fail string
		err := recover()
		got := strings.TrimSpace(out)
		want := "line 1\nline 2"

		if got != want {
			t.Errorf("Got: %#v, wanted: %#v.", got, want)
		}

		if err != nil {
			panic(err)
		}
	}()
}

func TestCleanLines(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
	)
	showColor = false
	screen = NewScreen(&showColor)
	screen.Display("a test", false, "do-not-delete")
	screen.InsertLine("", "delete")
	screen.Display(".", false, "delete")
	screen.InsertLine("another test", "do-not-delete")
	screen.InsertLine("", "delete")
	screen.CleanLines()
	numLines := screen.NumLines()
	if numLines != 3 {
		t.Errorf("Got: %d, wanted: %d.", numLines, 2)
	}
	screen.Clear()
}

func TestRemoveBlankLines(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
	)
	showColor = false
	screen = NewScreen(&showColor)
	screen.InsertLine("a test", "do-not-delete")
	screen.InsertLine("", "delete")
	screen.InsertLine(".", "delete")
	screen.InsertLine("another test", "do-not-delete")
	screen.InsertLine("", "delete")
	screen.RemoveBlankLines()
	numLines := screen.NumLines()
	if numLines != 3 {
		t.Errorf("Got: %d, wanted: %d.", numLines, 2)
	}
	//screen.Clear()
}

func TestClear(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
	)

	oldStdout := os.Stdout
	r, writeFile, err := os.Pipe()
	if err != nil {
		t.Errorf("FAIL")
	}

	os.Stdout = writeFile
	//
	outC := make(chan string)
	go func() {
		var buf strings.Builder
		_, err := io.Copy(&buf, r)
		_ = r.Close()
		if err != nil {
			t.Errorf("FAIL")
			os.Exit(1)
		}
		outC <- buf.String()
	}()

	//
	showColor = true
	screen = NewScreen(&showColor)

	screen.CarvePrint("line 1")
	screen.ShowPrint("line 2")
	screen.Display("line 3", true, "")
	screen.Clear()

	defer func() {
		// Close pipe, restore stdout, get output.
		writeFile.Close()
		os.Stdout = oldStdout
		out := <-outC

		//var fail string
		err := recover()
		got := strings.TrimSpace(out)
		want := "line 1\n\x1b[A\x1b[2Kline 1\nline 2\n\x1b[A\x1b[2K\x1b[A\x1b[2Kline 1\nline 3\n\x1b[A\x1b[2K\x1b[A\x1b[2K"

		if got != want {
			t.Errorf("Got: %#v, wanted: %#v.", got, want)
		}

		if err != nil {
			panic(err)
		}
	}()

}

func TestClearNamedLayers(t *testing.T) {
	var (
		screen    *Screen
		showColor bool
	)
	showColor = false
	screen = NewScreen(&showColor)
	screen.InsertLine("a test", "do-not-delete")
	screen.InsertLine("some test", "delete")
	screen.InsertLine("another test", "do-not-delete")
	screen.InsertLine("some more test", "delete")
	screen.ClearNamedLayers("delete")
	numLines := screen.NumLines()
	if numLines != 2 {
		t.Errorf("Got: %d, wanted: %d.", numLines, 2)
	}
	screen.Clear()
}

func TestColorOk(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is green"
	desiredTextWithColor = "\x1b[32mthis is green\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Ok(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Ok(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}
}

func TestColorWarning(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is yellow"
	desiredTextWithColor = "\x1b[33mthis is yellow\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Warning(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Warning(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}

func TestColorCritical(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is red"
	desiredTextWithColor = "\x1b[31mthis is red\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Critical(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Critical(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}

func TestColorBleu(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is bleu"
	desiredTextWithColor = "\x1b[34mthis is bleu\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Bleu(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Bleu(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}

func TestColorPurple(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is purple"
	desiredTextWithColor = "\x1b[35mthis is purple\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Purple(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Purple(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}

func TestColorTeal(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is teal"
	desiredTextWithColor = "\x1b[36mthis is teal\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Teal(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Teal(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}

func TestColorNormal(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen      *Screen
		showColor   bool
		resultText  string
		desiredText string
	)
	desiredText = "this is normal"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Normal(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Normal(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}

func TestColorEmphasis(t *testing.T) {
	// with noColor flag
	// without noColor flag
	var (
		screen               *Screen
		showColor            bool
		resultText           string
		desiredText          string
		desiredTextWithColor string
	)
	desiredText = "this is bold"
	desiredTextWithColor = "\x1b[1mthis is bold\x1b[0m"
	showColor = true
	screen = NewScreen(&showColor)
	resultText = screen.Emphasis(desiredText)
	if resultText != desiredTextWithColor {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredTextWithColor)
	}

	showColor = false
	screen = NewScreen(&showColor)
	resultText = screen.Emphasis(desiredText)
	if resultText != desiredText {
		t.Errorf("Got: %#v, wanted: %#v.", resultText, desiredText)
	}
}
