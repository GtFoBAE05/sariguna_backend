package entity

import "mime/multipart"

type ProductRepositoryInterface interface {
	CreateProduct(data ProductCore, image *multipart.FileHeader) error
	GetAllProduct() ([]ProductCore, error)
	GetProductById(id int) (ProductCore, error)
	GetProductByCategory(id int) ([]ProductCore, error)
	UpdateProduct(id int, data ProductCore, image *multipart.FileHeader) error
	DeleteProduct(id int) error
}

type ProductServiceInterface interface {
	CreateProduct(data ProductCore, image *multipart.FileHeader) error
	GetAllProduct() ([]ProductCore, error)
	GetProductById(id int) (ProductCore, error)
	GetProductByCategory(id int) ([]ProductCore, error)
	UpdateProduct(id int, data ProductCore, image *multipart.FileHeader) error
	DeleteProduct(id int) error
}
