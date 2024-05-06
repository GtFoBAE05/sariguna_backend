package route

import (
	"sariguna_backend/sariguna/app/user/handler"
	"sariguna_backend/sariguna/app/user/repository"
	"sariguna_backend/sariguna/app/user/service"
	"sariguna_backend/sariguna/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouteUser(r *gin.RouterGroup, db *sqlx.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := r.Group("/user")

	user.POST("/register", userHandler.Register)

	user.POST("/login", userHandler.Login)

	user.GET("/get-detail", middleware.Auth, userHandler.GetProfileById)

	user.PUT("/update-password", middleware.Auth, userHandler.UpdatePassword)

	user.GET("/pong", middleware.Auth, handler.CheckMiddleware)

}
