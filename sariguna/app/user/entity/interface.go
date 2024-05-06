package entity

type UserRepositoryInterface interface {
	Register(data UserCore) (UserCore, error)
	FindByEmail(email string) (UserCore, error)
	FindById(id string) (UserCore, error)
	UpdatePassword(id string, password string) error
}

type UserServiceInterface interface {
	Register(data UserCore) (UserCore, error)
	Login(email, password string) (UserCore, string, error)
	FindById(id string) (UserCore, error)
	UpdatePassword(id, password string) error
}
