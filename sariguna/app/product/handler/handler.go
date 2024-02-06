package handler

import (
	"net/http"
	"sariguna_backend/sariguna/app/product/dto/request"
	"sariguna_backend/sariguna/app/product/entity"
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

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_CREATE_DATA))
}

func (pch *ProductCategoryHandler) GetAllProductCategory(c *gin.Context) {

	res, errCreate := pch.productCategoryService.GetAllProductCategory()

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(201, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

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

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_UPDATE_DATA))
}

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

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_DELETE_DATA))
}
