package repository

import (
	"database/sql"
	"urlShortener/models"
	pb "urlShortener/proto/generated"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	CreateUrlShortener(urlShortener models.UrlShortener) (int, error)
	AllUrlShorteners() ([]*pb.UrlShortener, error)
	UpdateUrlShortener(urlShortener *pb.UrlShortener) error
	GetUrlShorteners(urlShorcut string) (*pb.UrlShortener, error)
	DeleteUrlShortener(urlID int64) error
}
