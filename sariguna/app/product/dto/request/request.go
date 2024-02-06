package request

type ProductCategoryCreate struct {
	CategoryName string `json:"CategoryName" binding:"required"`
}
