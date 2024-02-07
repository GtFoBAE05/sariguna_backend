package model

import "time"

type Product struct {
	Id           int       `db:"id"`
	CategoryId   int       `db:"category_id"`
	CategoryName string    `db:"category_name"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	Status       bool      `db:"status"`
	ImageUrl     string    `db:"image_url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
