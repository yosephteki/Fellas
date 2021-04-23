package Routes

import (
	"Fellas/Controllers"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func auth(c *gin.Context) {
	var tokenString string
	tokenStringRAW := c.Request.Header.Get("Authorization")
	if tokenStringRAW != "" {
		tokenStringSplit := strings.Split(tokenStringRAW, " ")
		tokenString = tokenStringSplit[1]
	} else {
		c.JSON(http.StatusUnauthorized, "token missing")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//this syntax to add group to endpoints
	grp1 := r.Group("/user-api")
	{
		grp1.GET("users", auth, Controllers.GetUsers)
		grp1.GET("user/:id", Controllers.GetUserById)
		grp1.POST("user", Controllers.CreateUser)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
		grp1.POST("userLogin", Controllers.Login)
	}
	grp2 := r.Group("/idea-api")
	{
		grp2.GET("ideas", Controllers.GetIdeas)
		grp2.GET("userIdea/:userId", Controllers.GetIdeaByUserId)
		grp2.GET("ideasUser/:userId", Controllers.GetUserIdea)
		grp2.GET("ideaUserJoin/:userId", Controllers.GetUserIdeaJoin)
	}
	return r
}
