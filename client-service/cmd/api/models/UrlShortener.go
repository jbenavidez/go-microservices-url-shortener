package models

import "time"

type UrlShortener struct {
	ID        int       `json:"id"`
	FullPath  string    `json:"full_path"`
	Shortcut  string    `json:"shortcut"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
