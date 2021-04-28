package Controllers

import (
	"Fellas/Models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user Models.User
	var userLogin Models.User
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&userLogin)
	if err != nil {
		fmt.Println("ERR MARSHAL", unmarshalErr)
	}
	err = Models.Login(&user, userLogin.Email, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		sign := jwt.New(jwt.GetSigningMethod("HS256"))
		token, err := sign.SignedString([]byte("secret"))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.JSON(http.StatusOK, token)
	}

}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserById(&user, id)
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
		c.JSON(http.StatusInternalServerError, "User not found")
	}
	// c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Fail to update user")
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
		c.JSON(http.StatusOK, userIdeas)
	}
}

func GetUserIdeaJoin(c *gin.Context) {
	userId := c.Params.ByName("userId")
	userIdeaJoins, err := Models.GetUserIdeaJoin(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userIdeaJoins)
	}

}
