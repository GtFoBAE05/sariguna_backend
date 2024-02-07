package request

import "sariguna_backend/sariguna/app/product/entity"

func ProductCreateToProductCore(data ProductCreate) entity.ProductCore {
	return entity.ProductCore{
		CategoryId:  data.ProductCategoryId,
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
	}
}
