package response

type ProductResponse struct {
	Id              int    `json:"id"`
	ProductCategory string `json:"productCategory"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ImageUrl        string `json:"imageUrl"`
}
