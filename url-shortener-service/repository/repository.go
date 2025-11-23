package repository

import "urlShortener/models"

type DatabaseRepo interface {
	AllUrlShortener() ([]*models.UrlShortener, error)
	CreateUrlShortener(urlShortener models.UrlShortener) (int, error)
}
