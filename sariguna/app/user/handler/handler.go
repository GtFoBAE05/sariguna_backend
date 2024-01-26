package handler

import (
	"sariguna_backend/sariguna/app/user/dto/request"
	"sariguna_backend/sariguna/app/user/dto/response"
	"sariguna_backend/sariguna/app/user/entity"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService entity.UserServiceInterface
}

func NewUserHandler(us entity.UserServiceInterface) *userHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) Register(c *gin.Context) {
	body := request.UserRegister{}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, gin.H{"error": "Invalid request payload"})
		return
	}

	request := request.UsersRequestRegisterToUsersCore(body)

	if _, errRegister := uh.userService.Register(request); errRegister != nil {
		c.JSON(500, gin.H{"error": errRegister})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully!", "data": response.UsersCoreToUsersRegisterResponse(request)})
}

func (uh *userHandler) Login(c *gin.Context) {
	body := request.UserLogin{}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, gin.H{"error": "Invalid request payload"})
		return
	}

	request := request.UsersRequestLoginToUsersCore(body)

	result, token, errLogin := uh.userService.Login(request.Email, request.Password)

	if errLogin != nil {
		c.JSON(500, gin.H{"error": errLogin.Error()})
		return
	}

	response := response.UsersCoreToLoginResponse(result, token)

	c.JSON(201, gin.H{"message": "User Login successfully!", "data": response})

}
