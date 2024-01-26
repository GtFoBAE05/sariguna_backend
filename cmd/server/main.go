package main

import (
	"fmt"
	"net/http"
	"sariguna_backend/sariguna/app/route"
	"sariguna_backend/sariguna/pkg/configs"
	"sariguna_backend/sariguna/pkg/database"
	"strings"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {
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

	app.Run()

}

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
