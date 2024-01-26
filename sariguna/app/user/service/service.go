package service

import (
	"errors"
	"sariguna_backend/sariguna/app/user/entity"
)

type userService struct {
	userRepo entity.UserRepositoryInterface
}

func NewUserService(userRepo entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(data entity.UserCore) (entity.UserCore, error) {
	//validasi

	dataUsers, err := us.userRepo.Register(data)
	if err != nil {
		return entity.UserCore{}, err
	}

	return dataUsers, nil

}

func (us *userService) Login(email, password string) (entity.UserCore, string, error) {
	dataUser, err := us.userRepo.FindByEmail(email)

	if err != nil {
		return entity.UserCore{}, "", err
	}

	if dataUser.Password != password {
		return entity.UserCore{}, "", errors.New("password tidak cocok")
	}

	return dataUser, "12345", nil
}
