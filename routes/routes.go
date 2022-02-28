package routes

import (
	"fmt"
	actions "gin/actions"
	middlewares "gin/middlewares"

	"github.com/gin-gonic/gin"
)

func Test() {
	fmt.Print("sample")
}

func DefineRoutes(router *gin.Engine) {

	/*
		API routes
	*/
	api := router.Group("api")
	{
		/**
		Albums
		*/
		albums := api.Group("albums")
		{
			albums.GET("/", actions.GetAlbums)
			albums.POST("/", actions.StoreAlbum)
			albums.GET("/:id", actions.GetAlbumByID)
		}
	}

	/*
		Web routes
	*/
	web := router.Group("web").Use(middlewares.Authenticate())
	{
		web.GET("/", actions.IndexWebPage)
	}
}
