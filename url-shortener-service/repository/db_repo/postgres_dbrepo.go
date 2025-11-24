package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"urlShortener/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllUrlShortener() ([]*models.UrlShortener, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			full_path, shortcut
		from
			url_shortener
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlShortenerList []*models.UrlShortener

	for rows.Next() {
		var urlShortener models.UrlShortener
		err := rows.Scan(
			&urlShortener.FullPath,
			&urlShortener.Shortcut,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		urlShortenerList = append(urlShortenerList, &urlShortener)
	}
	return urlShortenerList, nil
}

func (m *PostgresDBRepo) CreateUrlShortener(urlShortener models.UrlShortener) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
		insert into url_shortener
			(full_path, shortcut)
		values	
			($1,$2) 
		returning id
	`

	var newID int

	err := m.DB.QueryRowContext(ctx, stmt,
		urlShortener.FullPath,
		urlShortener.Shortcut,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}
	return newID, nil

}
