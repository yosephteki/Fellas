package Routes

import (
	"Fellas/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//this syntax to add group to endpoints
	grp1 := r.Group("/user-api")
	{
		grp1.GET("users", Controllers.GetUsers)
		grp1.GET("user/:id", Controllers.GetUserById)
		grp1.POST("user", Controllers.CreateUser)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}
