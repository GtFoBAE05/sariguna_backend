package entity

import "sariguna_backend/sariguna/app/user/model"

func UserCoreToUserModel(data UserCore) model.User {
	userModel := model.User{
		Id:       data.Id,
		Email:    data.Email,
		Password: data.Password,
		Fullname: data.Fullname,
		Role:     data.Role,
	}

	return userModel
}

func UserModelToUserCore(data model.User) UserCore {
	userCore := UserCore{
		Id:        data.Id,
		Email:     data.Email,
		Password:  data.Password,
		Fullname:  data.Fullname,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return userCore
}
