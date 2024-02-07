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

type ProductHandler struct {
	productService entity.ProductServiceInterface
}

func NewProductHandler(ps entity.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
	}
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	body := request.ProductCreate{}

	if errBind := c.ShouldBind(&body); errBind != nil {
		c.JSON(422, helpers.ErrorResponse(errBind.Error()))
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(422, helpers.ErrorResponse(err.Error()))
		return
	}

	request := request.ProductCreateToProductCore(body)

	if err := ph.productService.CreateProduct(request, file); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
		return
	}

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_CREATE_DATA))
}

func (ph *ProductHandler) GetAllProduct(c *gin.Context) {

	res, errCreate := ph.productService.GetAllProduct()

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(201, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

func (ph *ProductHandler) GetProductById(c *gin.Context) {

	productId := c.Param("id")
	idStr, errConv := strconv.Atoi(productId)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errConv.Error()))
		return
	}

	res, errCreate := ph.productService.GetProductById(idStr)

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(201, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

func (ph *ProductHandler) GetProductByCategory(c *gin.Context) {

	productId := c.Param("id")
	idStr, errConv := strconv.Atoi(productId)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errConv.Error()))
		return
	}

	res, errCreate := ph.productService.GetProductByCategory(idStr)

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(201, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	body := request.ProductCreate{}

	productId := c.Param("id")
	idStr, errConv := strconv.Atoi(productId)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errConv.Error()))
		return
	}

	if errBind := c.ShouldBind(&body); errBind != nil {
		c.JSON(422, errBind.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(422, helpers.ErrorResponse(constant.ERROR_INVALID_PAYLOAD))
		return
	}

	request := request.ProductCreateToProductCore(body)

	if err := ph.productService.UpdateProduct(idStr, request, file); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
		return
	}

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_UPDATE_DATA))
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {

	productId := c.Param("id")
	idStr, errConv := strconv.Atoi(productId)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errConv.Error()))
		return
	}

	if err := ph.productService.DeleteProduct(idStr); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
		return
	}

	c.JSON(201, helpers.SuccessResponse(constant.SUCCESS_DELETE_DATA))
}
