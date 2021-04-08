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
		grp1.GET("user", Controllers.GetUsers)

	}

	return r

}
