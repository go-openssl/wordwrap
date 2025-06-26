package wordwrap

import (
	"strings"
)

type Option func(*Wordwrap)

func WithWidth(width int) Option {
	return func(w *Wordwrap) {
		if width > 0 {
			w.width = width
		}
	}
}

func WithPrivateKeyBegin(begin string) Option {
	return func(w *Wordwrap) {
		strings.ReplaceAll(begin, "\n", "")
		begin += "\n"
		w.privateKeyBegin = begin
	}
}

func WithPrivateKeyEnd(end string) Option {
	return func(w *Wordwrap) {
		strings.ReplaceAll(end, "\n", "")
		end = "\n" + end
		w.privateKeyEnd = end
	}
}

func WithPublicKeyBegin(begin string) Option {
	return func(w *Wordwrap) {
		strings.ReplaceAll(begin, "\n", "")
		begin += "\n"
		w.publicKeyBegin = begin
	}
}

func WithPublicKeyEnd(end string) Option {
	return func(w *Wordwrap) {
		strings.ReplaceAll(end, "\n", "")
		end = "\n" + end
		w.publicKeyEnd = end
	}
}
