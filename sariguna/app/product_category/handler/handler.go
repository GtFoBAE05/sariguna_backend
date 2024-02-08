package handler

import (
	"net/http"
	"sariguna_backend/sariguna/app/product_category/dto/request"
	"sariguna_backend/sariguna/app/product_category/entity"
	"sariguna_backend/sariguna/pkg/constant"
	"sariguna_backend/sariguna/pkg/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
	productCategoryService entity.ProductCategoryServiceInterface
}

func NewProductCategoryHandler(pcs entity.ProductCategoryServiceInterface) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		productCategoryService: pcs,
	}
}

//	@Summary	Create product category
//	@Tags		product category
//	@Accept		json
//	@Produce	json
//	@Param		Create	body		request.ProductCategoryCreate	true	"Create"
//	@Success	201		{object}	helpers.SuccessResponseJson{}
//	@Failure	422		{object}	helpers.ErrorResponseJson{}
//	@Failure	400		{object}	helpers.ErrorResponseJson{}
//	@Router		/category/create [post]
func (pch *ProductCategoryHandler) CreateProductCategory(c *gin.Context) {
	body := request.ProductCategoryCreate{}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, helpers.ErrorResponse(constant.ERROR_INVALID_PAYLOAD))
		return
	}

	request := request.ProductCategoryCreateToProductCategoryCore(body)

	if errCreate := pch.productCategoryService.CreateProductCategory(request); errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(http.StatusCreated, helpers.SuccessResponse(constant.SUCCESS_CREATE_DATA))
}

//	@Summary	Get all product category
//	@Tags		product category
//	@Accept		json
//	@Produce	json
//
//	@Success	200	{object}	helpers.SuccessResponseJson{data=[]response.ProductCategoryResponse}
//	@Failure	422	{object}	helpers.ErrorResponseJson{}	"invalid payload"
//	@Failure	400	{object}	helpers.ErrorResponseJson{}
//	@Router		/category [get]
func (pch *ProductCategoryHandler) GetAllProductCategory(c *gin.Context) {

	res, errCreate := pch.productCategoryService.GetAllProductCategory()

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

//	@Summary	Update product category
//	@Tags		product category
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int								true	"Product Category Id"
//	@Param		Update	body		request.ProductCategoryCreate	true	"Update"
//	@Success	200		{object}	helpers.SuccessResponseJson{}
//	@Failure	422		{object}	helpers.ErrorResponseJson{}
//	@Failure	400		{object}	helpers.ErrorResponseJson{}
//	@Router		/category/update/{id} [put]
func (pch *ProductCategoryHandler) UpdateProductCategory(c *gin.Context) {
	body := request.ProductCategoryCreate{}
	id := c.Param("id")

	idStr, errConv := strconv.Atoi(id)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errConv.Error()))
		return
	}

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(422, helpers.ErrorResponse(constant.ERROR_INVALID_PAYLOAD))
		return
	}

	request := request.ProductCategoryCreateToProductCategoryCore(body)

	if errCreate := pch.productCategoryService.UpdateProductCategory(idStr, request); errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(constant.SUCCESS_UPDATE_DATA))
}

//	@Summary	Delete product category
//	@Tags		product category
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Product Category Id"
//	@Success	200	{object}	helpers.SuccessResponseJson{}
//	@Failure	422	{object}	helpers.ErrorResponseJson{}
//	@Failure	400	{object}	helpers.ErrorResponseJson{}
//	@Router		/category/delete/{id} [delete]
func (pch *ProductCategoryHandler) DeleteProductCategory(c *gin.Context) {
	id := c.Param("id")

	idStr, errConv := strconv.Atoi(id)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errConv.Error()))
		return
	}

	if errDelete := pch.productCategoryService.DeleteProductCategory(idStr); errDelete != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errDelete.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(constant.SUCCESS_DELETE_DATA))
}
