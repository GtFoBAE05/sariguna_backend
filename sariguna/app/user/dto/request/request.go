package request

type UserRegister struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
