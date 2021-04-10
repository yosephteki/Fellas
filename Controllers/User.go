package Controllers

import (
	"Fellas/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All User
func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id + "<---------------")
	var user Models.User
	err := Models.GetUserById(&user, id)
	fmt.Println(user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	var users []Models.User
	err := Models.GetAllUsers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS! CreateUser")
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS! UpdateUser")
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS! DeleteUser")
}
