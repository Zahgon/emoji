package main

import (
	"flag"
	"log"
	"os"
)

var pkgName string
var fileName string

func init() {
	log.SetFlags(log.Llongfile)

	flag.StringVar(&pkgName, "pkg", "emoji", "output package")
	flag.StringVar(&fileName, "o", "../../emoji_codemap.go", "output file")
	flag.Parse()
}

// TemplateData emoji_codemap.go template
type TemplateData struct {
	PkgName    string
	CodeMap    map[string]string
	RevCodeMap map[string][]string
}

const templateMapCode = `
package {{.PkgName}}

import (
	"sync"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// EMOJICODEMAP CODE GENERATION TOOL (github.com/kyokomi/emoji/cmd/generateEmojiCodeMap)
// DO NOT EDIT

var emojiCodeMap map[string]string
var emojiCodeMapInitOnce = sync.Once{}

func emojiCode() map[string]string {
	emojiCodeMapInitOnce.Do(func() {
		emojiCodeMap = map[string]string{
			{{range $key, $val := .CodeMap}}":{{$key}}:": {{$val}},
		{{end}}}
	})
	return emojiCodeMap
}

var emojiRevCodeMap map[string][]string
var emojiRevCodeMapInitOnce = sync.Once{}

func emojiRevCode() map[string][]string {
	emojiRevCodeMapInitOnce.Do(func() {
		emojiRevCodeMap = map[string][]string{
			{{range $key, $val := .RevCodeMap}} {{$key}}: { {{range $val}} ":{{.}}:", {{end}} },
		{{end}}}
	})
	return emojiRevCodeMap
}
`

func createCodeMap() (map[string]string, map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ensure deterministic ordering for aliases

func createCodeMapSource(pkgName string, emojiCodeMap map[string]string, emojiRevCodeMap map[string][]string) ([]byte, error) {
	_ = "STUB: not implemented"
	// Template GenerateSource
	return nil, nil
}

// gofmt

func main() {
	emojiCodeMap, emojiRevCodeMap, err := createCodeMap()
	if err != nil {
		log.Fatalln(err)
	}

	codeMapSource, err := createCodeMapSource(pkgName, emojiCodeMap, emojiRevCodeMap)
	if err != nil {
		log.Fatalln(err)
	}

	os.Remove(fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if _, err := file.Write(codeMapSource); err != nil {
		log.Fatalln(err)
	}
}
