package models

type Link struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
	Link  string `json:"link"`
}
