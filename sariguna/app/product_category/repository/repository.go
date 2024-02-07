package repository

import (
	"sariguna_backend/sariguna/app/product_category/entity"
	"sariguna_backend/sariguna/app/product_category/model"

	"github.com/jmoiron/sqlx"
)

type ProductCategoryRepository struct {
	db *sqlx.DB
}

func NewProductCategoryRepository(db *sqlx.DB) entity.ProductCategoryRepositoryInterface {
	return &ProductCategoryRepository{
		db: db,
	}
}

// CreateProductCategory implements entity.ProductCategoryRepositoryInterface.
func (pcr *ProductCategoryRepository) CreateProductCategory(data entity.ProductCategoryCore) error {
	request := entity.ProductCategoryCoreToProductCategoryModel(data)

	query := `INSERT INTO product_category (category_name) 
	VALUES (:category_name)`

	_, err := pcr.db.NamedExec(query, request)

	if err != nil {
		return err
	}

	return nil
}

// GetAllProductCategory implements entity.ProductCategoryRepositoryInterface.
func (pcr *ProductCategoryRepository) GetAllProductCategory() ([]entity.ProductCategoryCore, error) {
	productCategoryList := []model.ProductCategory{}

	query := `SELECT id, category_name from product_category`

	err := pcr.db.Select(&productCategoryList, query)
	if err != nil {
		return []entity.ProductCategoryCore{}, err
	}

	result := entity.ProductCategoryModelListToProductCategoryCoreList(productCategoryList)

	return result, nil
}

// UpdateProductCategory implements entity.ProductCategoryRepositoryInterface.
func (pcr *ProductCategoryRepository) UpdateProductCategory(id int, data entity.ProductCategoryCore) error {
	request := entity.ProductCategoryCoreToProductCategoryModel(data)

	query := `UPDATE product_category
	SET category_name = $1
	WHERE id = $2`

	_, err := pcr.db.Exec(query, request.CategoryName, id)

	if err != nil {
		return err
	}

	return nil
}

// DeleteProductCategory implements entity.ProductCategoryRepositoryInterface.
func (pcr *ProductCategoryRepository) DeleteProductCategory(id int) error {
	query := `DELETE FROM product_category
	WHERE id = $1`

	_, err := pcr.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
