package screen

type Line struct {
	content string
	keep    bool
	name    string // so we can delete layers based on name
}

func NewLine(content string, keep bool, name string) *Line {
	return &Line{
		content:content,
		keep: keep,
		name:name,
	}
}

func (line *Line) isBlank() bool {
	return line.content == ""
}
