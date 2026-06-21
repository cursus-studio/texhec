package internal

import (
	"engine/modules/console"
	"strings"
)

type service struct {
	permanent          string
	previousDrawnLines int
	toPrint            string
	print              func(string)
}

func NewConsole() console.Service {
	return &service{
		permanent:          "",
		previousDrawnLines: 0,
		print:              func(s string) { print(s) },
	}
}

func clearCurrentLine() string { return "\033[2K" }
func goToPreviousLine() string { return "\033[1A" }

func (console *service) PrintPermanent(text string) {
	console.permanent += text + "\n"
}

func (console *service) Print(text string) {
	console.toPrint += text
}

func (console *service) Flush() {
	var builder strings.Builder
	builder.WriteString(clearCurrentLine())
	for i := 0; i < console.previousDrawnLines; i++ {
		builder.WriteString(goToPreviousLine() + clearCurrentLine())
	}
	flushed := builder.String()
	console.print(flushed + console.permanent + console.toPrint)
	console.permanent = ""
	console.previousDrawnLines = strings.Count(console.toPrint, "\n")
	console.toPrint = ""
}
