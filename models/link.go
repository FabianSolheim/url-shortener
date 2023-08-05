package models

import (
	"encoding/json"
	"os"
)

type Link struct {
    Alias    string `json:"alias"`
    Link string `json:"link"`
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