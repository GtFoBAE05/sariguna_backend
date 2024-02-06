package route

import (
	"sariguna_backend/sariguna/app/product/handler"
	"sariguna_backend/sariguna/app/product/repository"
	"sariguna_backend/sariguna/app/product/service"
	"sariguna_backend/sariguna/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouteProductCategory(r *gin.RouterGroup, db *sqlx.DB) {
	productCategoryRepository := repository.NewProductCategoryRepository(db)
	productCategoryService := service.NewProductCategoryService(productCategoryRepository)
	productCategoryHandler := handler.NewProductCategoryHandler(productCategoryService)

	category := r.Group("/category")

	category.POST("/create", middleware.Auth, productCategoryHandler.CreateProductCategory)

	category.GET("", middleware.Auth, productCategoryHandler.GetAllProductCategory)

	category.PUT("/update/:id", middleware.Auth, productCategoryHandler.UpdateProductCategory)

	category.DELETE("/delete/:id", middleware.Auth, productCategoryHandler.DeleteProductCategory)

}
