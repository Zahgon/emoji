package main

const emojiDataJsonURL = "https://github.com/iamcal/emoji-data/raw/master/emoji.json"

// EmojiData json parse struct
type EmojiData struct {
	Unified     string `json:"unified"`
	ShortName   string `json:"short_name"`
	ObsoletedBy string `json:"obsoleted_by"`
}

// UnifiedToChar renders a character from its hexadecimal codepoint
func UnifiedToChar(unified string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func createEmojiDataCodeMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
