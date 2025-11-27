package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"urlShortener/models"
	pb "urlShortener/proto/generated"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllUrlShorteners() ([]*pb.UrlShortener, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, full_path, shortcut
		from
			url_shortener
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlShortenerList []*pb.UrlShortener

	for rows.Next() {
		var urlShortener pb.UrlShortener
		err := rows.Scan(
			&urlShortener.Id,
			&urlShortener.UrlPath,
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

func (m *PostgresDBRepo) UpdateUrlShortener(urlShortener *pb.UrlShortener) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		update url_shortener set full_path=$1, shortcut=$2
		where id = $3
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		urlShortener.UrlPath,
		urlShortener.Shortcut,
		urlShortener.Id,
	)
	if err != nil {
		return err
	}

	return nil

}

func (m *PostgresDBRepo) GetUrlShorteners(urlShorcut string) (*pb.UrlShortener, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	fmt.Println("loking form ", urlShorcut)
	query := `
		select
			id, full_path, shortcut
		from
			url_shortener
		where shortcut = $1
	`
	row := m.DB.QueryRowContext(ctx, query, urlShorcut)
	var urlShortener pb.UrlShortener

	err := row.Scan(
		&urlShortener.Id,
		&urlShortener.UrlPath,
		&urlShortener.Shortcut,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &urlShortener, nil

}
