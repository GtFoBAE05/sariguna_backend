package entity

import (
	"sariguna_backend/sariguna/app/product/model"
)

func ProductCategoryCoreToProductCategoryModel(data ProductCategoryCore) model.ProductCategory {
	return model.ProductCategory{
		CategoryName: data.CategoryName,
	}
}

func ProductCategoryModelToProductCategoryCore(data model.ProductCategory) ProductCategoryCore {
	return ProductCategoryCore{
		Id:           data.Id,
		CategoryName: data.CategoryName,
	}
}

func ProductCategoryCoreListToProductCategoryModelList(data []ProductCategoryCore) []model.ProductCategory {
	list := []model.ProductCategory{}
	for _, v := range data {
		result := ProductCategoryCoreToProductCategoryModel(v)
		list = append(list, result)
	}
	return list
}

func ProductCategoryModelListToProductCategoryCoreList(data []model.ProductCategory) []ProductCategoryCore {
	list := []ProductCategoryCore{}
	for _, v := range data {
		result := ProductCategoryModelToProductCategoryCore(v)
		list = append(list, result)
	}
	return list
}
