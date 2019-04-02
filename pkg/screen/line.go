package screen

// Line is a representation of a line on the Screen.
// `keep` is to indicate if the line should survive calls to clear the screen
// `name` is to give the line a name so to be able to clear it directly
type Line struct {
	content string
	keep    bool
	name    string
}

func NewLine(content string, keep bool, name string) *Line {
	return &Line{
		content: content,
		keep:    keep,
		name:    name,
	}
}

func (line *Line) isBlank() bool {
	emptyConds := []string{
		"",
		"\n",
		"\r",
	}
	for _, cond := range emptyConds {
		if line.content == cond {
			return true
		}
	}
	return false
}
