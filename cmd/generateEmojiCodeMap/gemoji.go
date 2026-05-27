package main

const gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"

// GemojiEmoji gemoji json parse struct
type GemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

func createGemojiCodeMap() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
