package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All User
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS!")
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS!")
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS!")
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS!")
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS!")
}
