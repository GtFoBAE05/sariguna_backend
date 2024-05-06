package repository

import (
	"mime/multipart"
	"sariguna_backend/sariguna/app/product/entity"
	"sariguna_backend/sariguna/app/product/model"
	"sariguna_backend/sariguna/pkg/images"
	"time"

	productcategory "sariguna_backend/sariguna/app/product_category/model"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) entity.ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

// CreateProduct implements entity.ProductRepositoryInterface.
func (pr *ProductRepository) CreateProduct(data entity.ProductCore, image *multipart.FileHeader) error {
	uploadUrl := "https://res.cloudinary.com/dxtawi0rz/image/upload/v1715013456/knnc6jdaaraitoik7gia.png"
	productCategory := productcategory.ProductCategory{}

	query := `SELECT id, category_name from product_category where id = $1`

	err := pr.db.Get(&productCategory, query, data.CategoryId)
	if err != nil {
		return err
	}
	if image != nil {
		uploadUrl, err = images.Upload(image, productCategory.CategoryName)
		if err != nil {
			return err
		}
		query = `INSERT INTO product (category_id, name, description, status, image_url) 
		VALUES ($1, $2, $3, $4, $5)`

		_, err = pr.db.Exec(query, data.CategoryId, data.Name, data.Description, data.Status, uploadUrl)

		if err != nil {
			return err
		}
	} else {
		query = `INSERT INTO product (category_id, name, description, status, image_url) 
		VALUES ($1, $2, $3, $4, $5)`

		_, err = pr.db.Exec(query, data.CategoryId, data.Name, data.Description, data.Status, uploadUrl)

		if err != nil {
			return err
		}

	}

	return nil
}

// DeleteProduct implements entity.ProductRepositoryInterface.
func (pr *ProductRepository) DeleteProduct(id int) error {
	res, err := pr.GetProductById(id)

	if err != nil {
		return err
	}

	if res.ImageUrl != "https://res.cloudinary.com/dxtawi0rz/image/upload/v1715013456/knnc6jdaaraitoik7gia.png" {
		errDelete := images.Remove(res.ImageUrl)

		if errDelete != nil {
			return errDelete
		}
	}

	query := `DELETE FROM product
	WHERE id = $1`

	_, err = pr.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

// GetAllProduct implements entity.ProductRepositoryInterface.
func (pr *ProductRepository) GetAllProduct() ([]entity.ProductCore, error) {
	productList := []model.Product{}

	query := `SELECT 
				p.id, 
				pc.id as category_id,
				pc.category_name as category_name, 
				p.name, p.description, 
				p.status, p.image_url, 
				p.created_at 
				FROM product p
				JOIN product_category pc
				ON p.category_id = pc.id`

	err := pr.db.Select(&productList, query)
	if err != nil {
		return []entity.ProductCore{}, err
	}

	result := entity.ProductModelListToProductCoreList(productList)

	return result, nil
}

// GetProductByCategory implements entity.ProductRepositoryInterface.
func (pr *ProductRepository) GetProductByCategory(id int) ([]entity.ProductCore, error) {
	productList := []model.Product{}

	query := `SELECT 
				p.id, 
				pc.category_name, 
				p.name, p.description, 
				p.status, p.image_url, 
				p.created_at 
				FROM product p
				JOIN product_category pc
				ON p.category_id = pc.id
				WHERE pc.id = $1
			`

	err := pr.db.Select(&productList, query, id)
	if err != nil {
		return []entity.ProductCore{}, err
	}

	result := entity.ProductModelListToProductCoreList(productList)

	return result, nil
}

// GetProductById implements entity.ProductRepositoryInterface.
func (pr *ProductRepository) GetProductById(id int) (entity.ProductCore, error) {
	product := model.Product{}
	query := `SELECT 
				p.id, 
				pc.category_name, 
				p.name, p.description, 
				p.status, p.image_url, 
				p.created_at 
				FROM product p
				JOIN product_category pc
				ON p.category_id = pc.id
				WHERE p.id = $1
			`

	err := pr.db.Get(&product, query, id)

	if err != nil {
		return entity.ProductCore{}, err
	}

	result := entity.ProductModelToProductCore(product)

	return result, nil
}

// UpdateProduct implements entity.ProductRepositoryInterface.
func (pr *ProductRepository) UpdateProduct(id int, data entity.ProductCore, image *multipart.FileHeader) error {
	request := entity.ProductCoreToProductModel(data)

	productCategory := productcategory.ProductCategory{}
	res, err := pr.GetProductById(id)

	if err != nil {
		return err
	}

	query := `SELECT id, category_name from product_category where id = $1`

	err = pr.db.Get(&productCategory, query, data.CategoryId)
	if err != nil {
		return err
	}

	if image != nil {

		if res.ImageUrl != "https://res.cloudinary.com/dxtawi0rz/image/upload/v1715013456/knnc6jdaaraitoik7gia.png" {
			errDelete := images.Remove(res.ImageUrl)

			if errDelete != nil {
				return errDelete
			}
		}
		newUploadUrl, err := images.Upload(image, productCategory.CategoryName)
		if err != nil {
			return err
		}

		query = `UPDATE product
			SET category_id = $1,
			name = $2,
			description = $3,
			status = $4,
			updated_at = $5,
			image_url = $6
			WHERE id = $7`

		_, err = pr.db.Exec(query, request.CategoryId, request.Name, request.Description, request.Status, time.Now(), newUploadUrl, id)

		if err != nil {
			return err
		}
	} else {
		query := `SELECT id, category_name from product_category where id = $1`

		err = pr.db.Get(&productCategory, query, data.CategoryId)
		if err != nil {
			return err
		}

		query = `UPDATE product
			SET category_id = $1,
			name = $2,
			description = $3,
			status = $4,
			updated_at = $5
			WHERE id = $6`

		_, err = pr.db.Exec(query, request.CategoryId, request.Name, request.Description, request.Status, time.Now(), id)

		if err != nil {
			return err
		}
	}

	return nil
}
