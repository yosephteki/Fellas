package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All User
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS! GetUser ")
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS! GetUsers")
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
