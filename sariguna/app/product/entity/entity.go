package entity

import "time"

type ProductCore struct {
	Id          int
	CategoryId  int
	Category    string
	Name        string
	Description string
	Status      bool
	ImageUrl    string
	CreatedAt   time.Time
}
