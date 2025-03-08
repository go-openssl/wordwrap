package wordwrap

import (
	"strings"
)

// convert converts the given text to a multi-line string with the specified width.
func convert(text string, width int) string {
	var lines []string
	runes := []rune(text)
	currentLine := ""

	for _, r := range runes {
		if len(currentLine) >= width {
			lines = append(lines, currentLine)
			currentLine = ""
		}
		currentLine += string(r)
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

type Wordwrap struct {
}

func New() *Wordwrap {
	return &Wordwrap{}
}

func (w *Wordwrap) ToPrivateKey(text string) string {
	str := "-----BEGIN RSA PRIVATE KEY-----\n"
	str += convert(text, 64)
	str += "\n-----END RSA PRIVATE KEY-----"

	return str
}

func (w *Wordwrap) ToPublicKey(text string) string {
	str := "-----BEGIN PUBLIC KEY-----\n"
	str += convert(text, 64)
	str += "\n-----END PUBLIC KEY-----"

	return str
}
