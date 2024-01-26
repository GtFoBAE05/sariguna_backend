package request

import (
	"sariguna_backend/sariguna/app/user/entity"

	"github.com/google/uuid"
)

func UsersRequestRegisterToUsersCore(data UserRegister) entity.UserCore {
	return entity.UserCore{
		Id:       uuid.New(),
		Fullname: data.FullName,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func UsersRequestLoginToUsersCore(data UserLogin) entity.UserCore {
	return entity.UserCore{
		Email:    data.Email,
		Password: data.Password,
	}
}
