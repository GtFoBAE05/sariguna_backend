package handler

import (
	"net/http"
	"sariguna_backend/sariguna/app/user/dto/request"
	"sariguna_backend/sariguna/app/user/dto/response"
	"sariguna_backend/sariguna/app/user/entity"
	"sariguna_backend/sariguna/pkg/constant"
	"sariguna_backend/sariguna/pkg/helpers"

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

//	@Summary		Register user
//	@Description	Register user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Register	body		request.UserRegister	true	"Register"
//	@Success		201			{object}	helpers.SuccessResponseJson{}
//	@Failure		422			{object}	helpers.ErrorResponseJson{}
//	@Failure		400			{object}	helpers.ErrorResponseJson{}
//	@Router			/user/register [post]
func (uh *userHandler) Register(c *gin.Context) {
	body := request.UserRegister{}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, helpers.ErrorResponse(constant.ERROR_INVALID_PAYLOAD))
		return
	}

	request := request.UsersRequestRegisterToUsersCore(body)

	if _, errRegister := uh.userService.Register(request); errRegister != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errRegister.Error()))
		return
	}

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_CREATE_DATA))
}

//	@Summary		Login user
//	@Description	Login user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Login	body		request.UserLogin	true	"Login"
//	@Success		200		{object}	helpers.SuccessResponseJson{data=response.UserRegisterResponse}
//	@Failure		422		{object}	helpers.ErrorResponseJson{}	"invalid payload"
//	@Failure		400		{object}	helpers.ErrorResponseJson{}
//	@Router			/user/login [post]
func (uh *userHandler) Login(c *gin.Context) {
	body := request.UserLogin{}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, helpers.ErrorResponse(constant.ERROR_INVALID_PAYLOAD))
		return
	}

	request := request.UsersRequestLoginToUsersCore(body)

	result, token, errLogin := uh.userService.Login(request.Email, request.Password)

	if errLogin != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errLogin.Error()))
		return
	}

	response := response.UsersCoreToLoginResponse(result, token)

	c.JSON(200, helpers.SuccessWithDataResponse(constant.SUCCESS_LOGIN, response))

}
