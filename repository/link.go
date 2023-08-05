package repository

import (
	"log"
	"url-shortener/db"
	"url-shortener/models"
)

func GetLink(alias string) (models.Link, error) {
	DB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	link := models.Link{}
	err = DB.Get(&link, "SELECT * FROM link WHERE alias=?", alias)
	if err != nil {
		return link, err
	}
	return link, nil
}

func CreateLink(link *models.Link) (models.Link, error) {
	DB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	stmt := "INSERT INTO link (alias, link) VALUES (?, ?)"
	_, err = DB.Exec(stmt, link.Alias, link.Link)
	if err != nil {
		return *link, err
	}
	return *link, nil

}
