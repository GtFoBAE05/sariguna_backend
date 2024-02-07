package route

import (
	"sariguna_backend/sariguna/app/product/handler"
	"sariguna_backend/sariguna/app/product/repository"
	"sariguna_backend/sariguna/app/product/service"
	"sariguna_backend/sariguna/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouteProduct(r *gin.RouterGroup, db *sqlx.DB) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	category := r.Group("/product")

	category.POST("/create", middleware.Auth, productHandler.CreateProduct)

	category.GET("", middleware.Auth, productHandler.GetAllProduct)

	category.GET("/byid/:id", middleware.Auth, productHandler.GetProductById)

	category.GET("/bycategory/:id", middleware.Auth, productHandler.GetProductByCategory)

	category.PUT("/update/:id", middleware.Auth, productHandler.UpdateProduct)

	category.DELETE("/delete/:id", middleware.Auth, productHandler.DeleteProduct)

}
