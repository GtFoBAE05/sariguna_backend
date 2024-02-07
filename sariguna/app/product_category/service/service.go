package service

import "sariguna_backend/sariguna/app/product_category/entity"

type ProductCategoryService struct {
	ProductCategoryRepository entity.ProductCategoryRepositoryInterface
}

func NewProductCategoryService(productCategoryRepository entity.ProductCategoryRepositoryInterface) entity.ProductCategoryServiceInterface {
	return &ProductCategoryService{
		ProductCategoryRepository: productCategoryRepository,
	}
}

// CreateProductCategory implements entity.ProductCategoryServiceInterface.
func (pcs *ProductCategoryService) CreateProductCategory(data entity.ProductCategoryCore) error {
	err := pcs.ProductCategoryRepository.CreateProductCategory(data)

	if err != nil {
		return err
	}

	return nil
}

// GetAllProductCategory implements entity.ProductCategoryServiceInterface.
func (pcs *ProductCategoryService) GetAllProductCategory() ([]entity.ProductCategoryCore, error) {
	result, err := pcs.ProductCategoryRepository.GetAllProductCategory()

	if err != nil {
		return []entity.ProductCategoryCore{}, err
	}

	return result, nil
}

// UpdateProductCategory implements entity.ProductCategoryServiceInterface.
func (pcs *ProductCategoryService) UpdateProductCategory(id int, data entity.ProductCategoryCore) error {
	err := pcs.ProductCategoryRepository.UpdateProductCategory(id, data)

	if err != nil {
		return err
	}

	return nil

}

// DeleteProductCategory implements entity.ProductCategoryServiceInterface.
func (pcs *ProductCategoryService) DeleteProductCategory(id int) error {
	err := pcs.ProductCategoryRepository.DeleteProductCategory(id)

	if err != nil {
		return err
	}

	return nil
}
