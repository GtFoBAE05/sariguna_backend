package model

type ProductCategory struct {
	Id           int    `db:"id"`
	CategoryName string `db:"category_name"`
}
