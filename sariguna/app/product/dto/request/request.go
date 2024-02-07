package request

import "mime/multipart"

type ProductCreate struct {
	ProductCategoryId int                   `form:"categoryid" binding:"required"`
	Image             *multipart.FileHeader `form:"image" binding:"required"`
	Name              string                `form:"name" binding:"required"`
	Description       string                `form:"description" binding:"required"`
	Status            bool                  `form:"status"`
}
