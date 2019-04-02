package screen

import "testing"

func TestLineWithEmptyStringIsBlank(t *testing.T) {
	var (
		line    *Line
		isBlank bool
	)
	line = NewLine("", true, "noname")
	isBlank = line.isBlank()
	if isBlank != true {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, true)
	}

	line = NewLine("", false, "noname")
	isBlank = line.isBlank()
	if isBlank != true {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, true)
	}

	line = NewLine("", false, "")
	isBlank = line.isBlank()
	if isBlank != true {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, true)
	}

	line = NewLine("something and more", true, "noname")
	isBlank = line.isBlank()
	if isBlank != false {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, false)
	}

	line = NewLine("something and more", false, "noname")
	isBlank = line.isBlank()
	if isBlank != false {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, false)
	}
}

func TestLineWithreturnCharIsBlank(t *testing.T) {
	var (
		line    *Line
		isBlank bool
	)

	line = NewLine("\n", false, "noname")
	isBlank = line.isBlank()
	if isBlank != true {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, true)
	}

	line = NewLine("\r", false, "noname")
	isBlank = line.isBlank()
	if isBlank != true {
		t.Errorf("line was not empty, got: %v, want: %v.", isBlank, true)
	}
}
