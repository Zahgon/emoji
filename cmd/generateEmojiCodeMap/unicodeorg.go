package main

import (
	"io"
)

const unicodeorgURL = "https://www.unicode.org/emoji/charts/emoji-list.html"

func createUnicodeorgMap() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

// UnicodeorgEmoji unicode.org emoji
type UnicodeorgEmoji struct {
	No            int
	Code          string
	ShortName     string
	OtherKeywords []string
}

var shortNameReplaces = []string{
	":", "",
	",", "",
	"⊛", "", // \U+229B
	"“", "", // \U+201C
	"”", "", // \U+201D
}

func generateUnicodeorgCodeMap(body io.ReadCloser) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
