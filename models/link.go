package models

import (
	"encoding/json"
	"os"
)

type Link struct {
    ShortLink    string `json:"shortLink"`
    OriginalLink string `json:"originalLink"`
}


func GetLinks() ([]Link, error) {
	result, err := os.ReadFile("links.json")
	if(err != nil) {
		return nil, err
	}
	var links []Link
	json.Unmarshal(result, &links)
	return links, nil
}