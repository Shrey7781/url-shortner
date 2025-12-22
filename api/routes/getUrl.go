package routes

import (
	"net/http"

	"github.com/Shrey7781/url-shortner/api/database"
	"github.com/gin-gonic/gin"
)

func GetByShortID(c *gin.Context) {
	shortID := c.Param("shortID")
	
	val, err := database.Client.Get(database.Ctx, shortID).Result()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found for given ShortID",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": val,
	})
}
