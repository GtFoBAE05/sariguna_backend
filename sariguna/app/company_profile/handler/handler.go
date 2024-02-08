package handler

import (
	"net/http"
	"sariguna_backend/sariguna/app/company_profile/dto/request"
	"sariguna_backend/sariguna/app/company_profile/entity"
	"sariguna_backend/sariguna/pkg/constant"
	"sariguna_backend/sariguna/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type CompanyProfileHandler struct {
	CompanyProfileService entity.CompanyProfileServiceInterface
}

func NewCompanyProfileHandler(cps entity.CompanyProfileServiceInterface) *CompanyProfileHandler {
	return &CompanyProfileHandler{
		CompanyProfileService: cps,
	}
}

// @Summary	Get company profile
// @Tags		company profile
// @Accept		json
// @Produce	json
// @Success	200	{object}	helpers.SuccessResponseJson{data=response.CompanyProfileResponse}
// @Failure	422	{object}	helpers.ErrorResponseJson{}
// @Failure	400	{object}	helpers.ErrorResponseJson{}
// @Router		/profile/ [get]
func (cph *CompanyProfileHandler) GetCompanyProfile(c *gin.Context) {
	res, err := cph.CompanyProfileService.GetCompanyProfile()

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

// @Summary	Update company profile
// @Tags		company profile
// @Accept		json
// @Produce	json
// @Param		Update	body		request.CompanyProfileUpdate	true	"Update"
// @Success	200		{object}	helpers.SuccessResponseJson{}
// @Failure	422		{object}	helpers.ErrorResponseJson{}
// @Failure	400		{object}	helpers.ErrorResponseJson{}
// @Router		/profile/update [put]
func (cph *CompanyProfileHandler) UpdateCompanyProfile(c *gin.Context) {
	body := request.CompanyProfileUpdate{}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, helpers.ErrorResponse(constant.ERROR_INVALID_PAYLOAD))
		return
	}

	request := request.CompanyProfileCreateToCompanyProfileCore(body)

	err := cph.CompanyProfileService.UpdateCompanyProfile(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(constant.SUCCESS_UPDATE_DATA))
}
