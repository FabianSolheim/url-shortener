package repository

import (
	"url-shortener/models"

	"github.com/jmoiron/sqlx"
)

type LinkRepository struct {
	Conn *sqlx.DB
}

func NewLinkRepository(conn *sqlx.DB) *LinkRepository {
	return &LinkRepository{Conn: conn}
}


func (s *LinkRepository) GetOneLink(alias string) (models.Link, error) {
	link := models.Link{}
	err := s.Conn.Get(&link, "SELECT alias, link, id FROM link WHERE alias=?", alias)
	if err != nil {
		return link, err
	}
	return link, nil
}

func (s *LinkRepository) CreateLink(link *models.Link) (int, error) {
	stmt := "INSERT INTO link (alias, link) VALUES (?, ?)"
	res, err := s.Conn.Exec(stmt, link.Alias, link.Link)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}
