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
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserById(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func GetUserIdea(c *gin.Context) {
	userId := c.Params.ByName("userId")
	userIdeas, err := Models.GetUserIdeas(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println(userIdeas)
		c.JSON(http.StatusOK, userIdeas)
	}
}
