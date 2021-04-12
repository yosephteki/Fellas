package Controllers

import (
	"Fellas/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdeas(c *gin.Context) {
	var ideas []Models.Idea
	err := Models.GetAllIdeas(&ideas)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ideas)
	}
}

func GetIdeaByUserId(c *gin.Context) {
	id := c.Params.ByName("userId")
	var idea Models.Idea
	err := Models.GetIdeaByUserId(&idea, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, idea)
	}
}
