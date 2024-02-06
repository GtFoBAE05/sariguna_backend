package request

import "sariguna_backend/sariguna/app/product/entity"

func ProductCategoryCreateToProductCategoryCore(data ProductCategoryCreate) entity.ProductCategoryCore {
	return entity.ProductCategoryCore{
		CategoryName: data.CategoryName,
	}
}
