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
	tokenStringSplit := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(tokenStringSplit) > 1 {
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
		grp1.GET("userFriends/:id", auth, Controllers.GetUserFriends)
		grp1.GET("users", auth, Controllers.GetUsers)
		grp1.GET("user/:id", auth, Controllers.GetUserById)
		grp1.POST("user", Controllers.CreateUser)
		grp1.PUT("user/:id", auth, Controllers.UpdateUser)
		grp1.DELETE("user/:id", auth, Controllers.DeleteUser)
		grp1.POST("userLogin", Controllers.Login)
	}
	grp2 := r.Group("/idea-api")
	{
		grp2.GET("ideas", auth, Controllers.GetIdeas)
		grp2.GET("userIdea/:userId", auth, Controllers.GetIdeaByUserId)
		grp2.GET("ideasUser/:userId", auth, Controllers.GetUserIdea)
		grp2.GET("ideaUserJoin/:userId", auth, Controllers.GetUserIdeaJoin)
	}
	return r
}
