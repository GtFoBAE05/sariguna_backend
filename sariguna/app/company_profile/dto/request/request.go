package request

type CompanyProfileUpdate struct {
	Sejarah string `json:"sejarah" binding:"required"`
	Visi    string `json:"visi" binding:"required"`
	Misi    string `json:"misi" binding:"required"`
}
