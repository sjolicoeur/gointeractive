package formating

import (
	. "github.com/logrusorgru/aurora"
)

func Warning(text string, au Aurora) string {
	return Sprintf(au.Brown(text))
}


func Critical(text string, au Aurora) string {
	return Sprintf(au.Red(text))
}

func Emphasis(text string, au Aurora) string {
	return Sprintf(au.Bold(text))
}

func Invert(text string, au Aurora) string {
	return Sprintf(au.Inverse(text))
}

func Normal(text string, au Aurora) string {
	return text
}


func Ok(text string, au Aurora) string {
	return Sprintf(au.Green(text))
}

func Bleu(text string, au Aurora) string {
	return Sprintf(au.Blue(text))
}

func Teal(text string, au Aurora) string {
	return Sprintf(au.Cyan(text))
}

func Purple(text string, au Aurora) string {
	return Sprintf(au.Magenta(text))
}



