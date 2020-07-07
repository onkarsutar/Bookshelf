package model

import "time"

type Book struct {
	ID          string    `json:"id"`
	Author      string    `json:"author" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	ISBN        string    `json:"isbn" validate:"required"`
	ReleaseDate time.Time `json:"releaseDate"`
}
