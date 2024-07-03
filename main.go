package main

import (
	"fmt"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"skripsi.id/obfuss/libs"
	"skripsi.id/obfuss/routes"
)

// @title			Obfuss API
// @version		1
// @description	Obfuss API Documentation
// @host			localhost:6996
// @BasePath		/api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	initEnv()
	logger := libs.NewLogger()
	db := libs.NewMysqlConn(logger)

	app := gin.Default()

	app.Use(helmet.Default())
	api := app.Group("/api")

	routes.Setup(api, db, logger)

	app.Static("/public", "./public")

	app.Run(fmt.Sprint(":", os.Getenv("PORT")))
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	fmt.Println("env loaded")
}
