package api

import (
	"fmt"
	"net/http"
	"sariguna_backend/sariguna/app/route"
	"sariguna_backend/sariguna/pkg/configs"
	"sariguna_backend/sariguna/pkg/database"
	"strings"

	_ "sariguna_backend/docs"

	"github.com/gin-gonic/gin"
)

// swag init -g ./api/index.go
var app *gin.Engine

// @title		Swagger Exampaaale API
// @host		https://sariguna-backend.vercel.app
// @BasePath	/api
// @schemes	http
func init() {
	// gin.SetMode(gin.ReleaseMode)

	app = gin.New()

	app.Use(CORSMiddleware())

	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	cfg := configs.InitConfig()
	db, err := configs.ConnectPostgres(cfg)

	if err != nil {
		panic(err)
	}

	err = database.Migration(db)

	if err != nil {
		panic(err)
	}

	rg := app.Group("/api")

	route.SetupRoute(rg, db)

	app.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

// func main() {
// 	app.Run(":8080")
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
