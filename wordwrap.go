package wordwrap

import (
	"strings"
)

// Wrap wraps a string into multiple lines, with a given width.
func Wrap(text string, width int) string {
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

func ToSslText(text string) string {
	str := "-----BEGIN RSA PRIVATE KEY-----\n"
	str += Wrap(text, 64)
	str += "\n-----END RSA PRIVATE KEY-----"

	return str
}
