package repository

import (
	"database/sql"
	"urlShortener/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllUrlShortener() ([]*models.UrlShortener, error)
	CreateUrlShortener(urlShortener models.UrlShortener) (int, error)
}
