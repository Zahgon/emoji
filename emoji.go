// Package emoji terminal output.
package emoji

import (
	"bytes"
	"io"
	"regexp"
)

//go:generate generateEmojiCodeMap -pkg emoji -o emoji_codemap.go

// Replace Padding character for emoji.
var (
	ReplacePadding = " "
)

// CodeMap gets the underlying map of emoji.
func CodeMap() map[string]string {
	_ = "STUB: not implemented"

	// RevCodeMap gets the underlying map of emoji.
	return nil
}

func RevCodeMap() map[string][]string { _ = "STUB: not implemented"; return nil }

func AliasList(shortCode string) []string { _ = "STUB: not implemented"; return nil }

// HasAlias flags if the given `shortCode` has multiple aliases with other
// codes.
func HasAlias(shortCode string) bool { _ = "STUB: not implemented"; return false }

// NormalizeShortCode normalizes a given `shortCode` to a deterministic alias.
func NormalizeShortCode(shortCode string) string { _ = "STUB: not implemented"; return "" }

// regular expression that matches :flag-[countrycode]:
var flagRegexp = regexp.MustCompile(":flag-([a-z]{2}):")

// Emojize Converts the string passed as an argument to a emoji. For unsupported emoji, the string passed as an argument is returned as is.
func Emojize(x string) string { _ = "STUB: not implemented"; return "" }

// regionalIndicator maps a lowercase letter to a unicode regional indicator
func regionalIndicator(i byte) string { _ = "STUB: not implemented"; return "" }

func replaceEmoji(input *bytes.Buffer) string { _ = "STUB: not implemented"; return "" }

// not replace

func compile(x string) string { _ = "STUB: not implemented"; return "" }

// Print is fmt.Print which supports emoji
func Print(a ...interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Println is fmt.Println which supports emoji
func Println(a ...interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Printf is fmt.Printf which supports emoji
func Printf(format string, a ...interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Fprint is fmt.Fprint which supports emoji
func Fprint(w io.Writer, a ...interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Fprintln is fmt.Fprintln which supports emoji
func Fprintln(w io.Writer, a ...interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Fprintf is fmt.Fprintf which supports emoji
func Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Sprint is fmt.Sprint which supports emoji
func Sprint(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Sprintf is fmt.Sprintf which supports emoji
func Sprintf(format string, a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Errorf is fmt.Errorf which supports emoji
func Errorf(format string, a ...interface{}) error { _ = "STUB: not implemented"; return nil }
