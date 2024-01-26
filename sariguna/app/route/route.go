package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoute(r *gin.RouterGroup, db *sqlx.DB) {

	RouteUser(r, db)

}
