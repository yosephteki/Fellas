package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All User
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS!")
}
