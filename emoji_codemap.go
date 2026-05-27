package emoji

import (
	"sync"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// EMOJICODEMAP CODE GENERATION TOOL (github.com/kyokomi/emoji/cmd/generateEmojiCodeMap)
// DO NOT EDIT

var emojiCodeMap map[string]string
var emojiCodeMapInitOnce = sync.Once{}

func emojiCode() map[string]string { _ = "STUB: not implemented"; return nil }

var emojiRevCodeMap map[string][]string
var emojiRevCodeMapInitOnce = sync.Once{}

func emojiRevCode() map[string][]string { _ = "STUB: not implemented"; return nil }
