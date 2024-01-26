package entity

type UserRepositoryInterface interface {
	Register(data UserCore) (UserCore, error)
	FindByEmail(email string) (UserCore, error)
}

type UserServiceInterface interface {
	Register(data UserCore) (UserCore, error)
	Login(email, password string) (UserCore, string, error)
}
