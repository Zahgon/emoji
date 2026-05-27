package main

const emojoV5DBJsonURL = "https://raw.githubusercontent.com/CodeFreezr/emojo/master/db/v5/emoji-v5.json"

// Emojo json parse struct
type Emojo struct {
	No          int    `json:"No"`
	Emoji       string `json:"Emoji"`
	Category    string `json:"Category"`
	SubCategory string `json:"SubCategory"`
	Unicode     string `json:"Unicode"`
	Name        string `json:"Name"`
	Tags        string `json:"Tags"`
	Shortcode   string `json:"Shortcode"`
}

func createEmojoCodeMap() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
