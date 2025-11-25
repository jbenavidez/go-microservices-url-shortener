package repository

import (
	"database/sql"
	"urlShortener/models"
	pb "urlShortener/proto"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	CreateUrlShortener(urlShortener models.UrlShortener) (int, error)
	AllUrlShorteners() ([]*pb.UrlShortener, error)
}
