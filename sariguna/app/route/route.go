package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "sariguna_backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoute(r *gin.RouterGroup, db *sqlx.DB) {

	r.GET("/docs/*any", FixRequestUri)
	RouteUser(r, db)
	RouteProductCategory(r, db)

}

func FixRequestUri(c *gin.Context) {
	if c.Request.RequestURI == "" {
		c.Request.RequestURI = c.Request.URL.Path
	}

	ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
}
