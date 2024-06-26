package request

import "mime/multipart"

type ProductCreate struct {
	ProductCategoryId int                   `form:"categoryid" binding:"required"`
	Image             *multipart.FileHeader `form:"image" swaggerignore:"true"`
	Name              string                `form:"name" binding:"required"`
	Description       string                `form:"description" binding:"required"`
	Status            bool                  `form:"status"`
}
