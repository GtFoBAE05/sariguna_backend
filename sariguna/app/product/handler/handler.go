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

//	@Summary	Create product
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		Update	formData	request.ProductCreate	true	"Update"
//	@Param		Update	formData	file					true	"image"
//	@Success	201		{object}	helpers.SuccessResponseJson{}
//	@Failure	422		{object}	helpers.ErrorResponseJson{}
//	@Failure	400		{object}	helpers.ErrorResponseJson{}
//	@Router		/product/create [post]
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

//	@Summary	Get all product
//	@Tags		product
//	@Accept		json
//	@Produce	json
//
//	@Success	200	{object}	helpers.SuccessResponseJson{data=[]response.ProductResponse}
//	@Failure	422	{object}	helpers.ErrorResponseJson{}	"invalid payload"
//	@Failure	400	{object}	helpers.ErrorResponseJson{}
//	@Router		/product [get]
func (ph *ProductHandler) GetAllProduct(c *gin.Context) {

	res, errCreate := ph.productService.GetAllProduct()

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(errCreate.Error()))
		return
	}

	c.JSON(201, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

//	@Summary	Get product by product id
//	@Tags		product
//	@Accept		json
//	@Produce	json
//
//	@Param		id	path		int	true	"Product Id"
//	@Success	200	{object}	helpers.SuccessResponseJson{data=response.ProductResponse}
//	@Failure	422	{object}	helpers.ErrorResponseJson{}	"invalid payload"
//	@Failure	400	{object}	helpers.ErrorResponseJson{}
//	@Router		/product/byid/{id} [get]
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

	c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

//	@Summary	Get product by category id
//	@Tags		product
//	@Accept		json
//	@Produce	json
//
//	@Param		id	path		int	true	"category Id"
//	@Success	200	{object}	helpers.SuccessResponseJson{data=[]response.ProductResponse}
//	@Failure	422	{object}	helpers.ErrorResponseJson{}	"invalid payload"
//	@Failure	400	{object}	helpers.ErrorResponseJson{}
//	@Router		/product/bycategory/{id} [get]
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

	c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(constant.SUCCESS_GET_DATA, res))
}

//	@Summary	Update product
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int						true	"Product Id"
//	@Param		Update	formData	request.ProductCreate	true	"Update"
//	@Param		Update	formData	file					true	"image"
//	@Success	200		{object}	helpers.SuccessResponseJson{}
//	@Failure	422		{object}	helpers.ErrorResponseJson{}
//	@Failure	400		{object}	helpers.ErrorResponseJson{}
//	@Router		/product/update/{id} [put]
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

	c.JSON(http.StatusOK, helpers.SuccessResponse(constant.SUCCESS_UPDATE_DATA))
}

//	@Summary	Delete product
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Product Id"
//	@Success	200	{object}	helpers.SuccessResponseJson{}
//	@Failure	422	{object}	helpers.ErrorResponseJson{}
//	@Failure	400	{object}	helpers.ErrorResponseJson{}
//	@Router		/product/delete/{id} [delete]
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

	c.JSON(http.StatusOK, helpers.SuccessResponse(constant.SUCCESS_DELETE_DATA))
}
