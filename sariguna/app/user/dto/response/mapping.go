package response

import "sariguna_backend/sariguna/app/user/entity"

func UsersCoreToUsersRegisterResponse(data entity.UserCore) UserRegisterResponse {
	return UserRegisterResponse{
		Id:       data.Id.String(),
		Fullname: data.Fullname,
		Email:    data.Email,
	}
}

func UsersCoreToLoginResponse(data entity.UserCore, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:       data.Id.String(),
		Fullname: data.Fullname,
		Email:    data.Email,
		Token:    token,
	}
}
