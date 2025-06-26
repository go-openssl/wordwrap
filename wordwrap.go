package wordwrap

import (
	"strings"
)

// convert converts the given text to a multi-line string with the specified width.
func convert(text string, width int) string {
	strings.ReplaceAll(text, "\n", "")
	strings.ReplaceAll(text, "\r", "")
	strings.ReplaceAll(text, "\t", "")
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
	width           int
	privateKeyBegin string
	privateKeyEnd   string
	publicKeyBegin  string
	publicKeyEnd    string
}

func New(opts ...Option) *Wordwrap {
	w := &Wordwrap{
		width:           64,
		privateKeyBegin: "-----BEGIN RSA PRIVATE KEY-----\n",
		privateKeyEnd:   "\n-----END RSA PRIVATE KEY-----",
		publicKeyBegin:  "-----BEGIN PUBLIC KEY-----\n",
		publicKeyEnd:    "\n-----END PUBLIC KEY-----",
	}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

func (w *Wordwrap) ToPrivateKey(text string) string {
	str := w.privateKeyBegin
	str += convert(text, w.width)
	str += w.privateKeyEnd
	return str
}

func (w *Wordwrap) ToPublicKey(text string) string {
	str := w.publicKeyBegin
	str += convert(text, w.width)
	str += w.publicKeyEnd
	return str
}
