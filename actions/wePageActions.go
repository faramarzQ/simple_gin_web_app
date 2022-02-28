package actions

import (
	models "gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexWebPage(c *gin.Context) {

	albums := models.Albums

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"albums": albums,
			"title":  "Home Page",
		},
	)
}
