package response

import "sariguna_backend/sariguna/app/product/entity"

func ProductCoreToProductResponse(data entity.ProductCore) ProductResponse {
	return ProductResponse{
		Id:              data.Id,
		ProductCategory: data.Category,
		Name:            data.Name,
		Description:     data.Description,
		ImageUrl:        data.ImageUrl,
	}
}

func ProductListToProductListResponse(data []entity.ProductCore) []ProductResponse {
	list := []ProductResponse{}
	for _, v := range data {
		result := ProductCoreToProductResponse(v)
		list = append(list, result)
	}
	return list
}
