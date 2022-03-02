package main

import (
	DB "gin/database"
	h "gin/helpers"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {

}

func init() {
	bootstrapDatabase()

	startServer()
}

func bootstrapDatabase() {
	DB.Con()
}

func startServer() {
	router := gin.Default()

	routes.DefineRoutes(router)

	router.LoadHTMLGlob("templates/*")

	router.Run(":" + h.ENV("APP_PORT"))
}
