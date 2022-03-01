package main

import (
	h "gin/helpers"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {

}

func init() {
	startServer()
}

func startServer() {
	router := gin.Default()

	routes.DefineRoutes(router)

	router.LoadHTMLGlob("templates/*")

	router.Run(":" + h.ENV("APP_PORT"))
}
