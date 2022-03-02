package actions

import (
	DB "gin/database"
	m "gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []m.User
	DB.Con().Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {

	var user m.User
	DB.Con().First(&user, c.Param("id"))

	c.IndentedJSON(http.StatusOK, &user)
}

func StoreUser(c *gin.Context) {

	var newUser m.User

	if err := c.BindJSON(&newUser); err != nil {
		panic(err.Error())
	}

	DB.Con().Create(newUser)
}
