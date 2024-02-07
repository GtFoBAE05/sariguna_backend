package entity

import "sariguna_backend/sariguna/app/product/model"

func ProductCoreToProductModel(data ProductCore) model.Product {
	return model.Product{
		Id:           data.Id,
		CategoryId:   data.CategoryId,
		CategoryName: data.Category,
		Name:         data.Name,
		Description:  data.Description,
		Status:       data.Status,
		ImageUrl:     data.ImageUrl,
	}
}

func ProductModelToProductCore(data model.Product) ProductCore {
	return ProductCore{
		Id:          data.Id,
		CategoryId:  data.CategoryId,
		Category:    data.CategoryName,
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
		ImageUrl:    data.ImageUrl,
		CreatedAt:   data.CreatedAt,
	}
}

func ProductCoreListToProductModelList(data []ProductCore) []model.Product {
	list := []model.Product{}
	for _, v := range data {
		result := ProductCoreToProductModel(v)
		list = append(list, result)
	}
	return list
}

func ProductModelListToProductCoreList(data []model.Product) []ProductCore {
	list := []ProductCore{}
	for _, v := range data {
		result := ProductModelToProductCore(v)
		list = append(list, result)
	}
	return list
}
