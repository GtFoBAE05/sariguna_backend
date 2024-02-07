package service

import (
	"mime/multipart"
	"sariguna_backend/sariguna/app/product/entity"
)

type ProductService struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewProductService(productRepository entity.ProductRepositoryInterface) entity.ProductServiceInterface {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

// CreateProduct implements entity.ProductServiceInterface.
func (ps *ProductService) CreateProduct(data entity.ProductCore, image *multipart.FileHeader) error {
	err := ps.ProductRepository.CreateProduct(data, image)

	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct implements entity.ProductServiceInterface.
func (ps *ProductService) DeleteProduct(id int) error {
	err := ps.ProductRepository.DeleteProduct(id)

	if err != nil {
		return err
	}

	return nil
}

// GetAllProduct implements entity.ProductServiceInterface.
func (ps *ProductService) GetAllProduct() ([]entity.ProductCore, error) {
	result, err := ps.ProductRepository.GetAllProduct()

	if err != nil {
		return []entity.ProductCore{}, err
	}

	return result, nil
}

// GetProductByCategory implements entity.ProductServiceInterface.
func (ps *ProductService) GetProductByCategory(id int) ([]entity.ProductCore, error) {
	result, err := ps.ProductRepository.GetProductByCategory(id)

	if err != nil {
		return []entity.ProductCore{}, err
	}

	return result, nil
}

// GetProductById implements entity.ProductServiceInterface.
func (ps *ProductService) GetProductById(id int) (entity.ProductCore, error) {
	result, err := ps.ProductRepository.GetProductById(id)

	if err != nil {
		return entity.ProductCore{}, err
	}

	return result, nil
}

// UpdateProduct implements entity.ProductServiceInterface.
func (ps *ProductService) UpdateProduct(id int, data entity.ProductCore, image *multipart.FileHeader) error {
	err := ps.ProductRepository.UpdateProduct(id, data, image)

	if err != nil {
		return err
	}

	return nil
}
