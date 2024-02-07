package response

import "sariguna_backend/sariguna/app/product_category/entity"

func ProductCategoryCoreToProductCategoryResponse(data entity.ProductCategoryCore) ProductCategoryResponse {
	return ProductCategoryResponse{
		Id:                  data.Id,
		ProductCategoryName: data.CategoryName,
	}
}

func ProductCategoryListCoreToProductCategoryListResponse(data []entity.ProductCategoryCore) []ProductCategoryResponse {
	list := []ProductCategoryResponse{}
	for _, v := range data {
		result := ProductCategoryCoreToProductCategoryResponse(v)
		list = append(list, result)
	}
	return list
}
