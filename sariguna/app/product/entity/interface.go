package entity

type ProductCategoryRepositoryInterface interface {
	CreateProductCategory(data ProductCategoryCore) error
	GetAllProductCategory() ([]ProductCategoryCore, error)
	UpdateProductCategory(id int, data ProductCategoryCore) error
	DeleteProductCategory(id int) error
}

type ProductCategoryServiceInterface interface {
	CreateProductCategory(data ProductCategoryCore) error
	GetAllProductCategory() ([]ProductCategoryCore, error)
	UpdateProductCategory(id int, data ProductCategoryCore) error
	DeleteProductCategory(id int) error
}
