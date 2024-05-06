package service

import (
	"errors"
	"sariguna_backend/sariguna/app/user/entity"
	"sariguna_backend/sariguna/pkg/helpers"
	"sariguna_backend/sariguna/pkg/jwt"
)

type userService struct {
	userRepo entity.UserRepositoryInterface
}

func NewUserService(userRepo entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &userService{
		userRepo: userRepo,
	}
}

// FindById implements entity.UserServiceInterface.
func (us *userService) FindById(id string) (entity.UserCore, error) {
	dataUser, err := us.userRepo.FindById(id)

	if err != nil {
		return entity.UserCore{}, err
	}

	return dataUser, nil
}

// UpdatePassword implements entity.UserServiceInterface.
func (us *userService) UpdatePassword(id string, password string) error {
	err := us.userRepo.UpdatePassword(id, password)
	if err != nil {
		return err
	}

	return nil
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

	if !helpers.CompareHash(dataUser.Password, password) {
		return entity.UserCore{}, "", errors.New("incorrect password")
	}

	token, err := jwt.GenerateToken(dataUser.Id.String(), dataUser.Role)

	if err != nil {
		return entity.UserCore{}, "", err
	}

	return dataUser, token, nil
}
