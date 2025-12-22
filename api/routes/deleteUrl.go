package routes

import (
	"net/http"

	"github.com/Shrey7781/url-shortner/api/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")
	val, err := database.Client.Get(database.Ctx, shortID).Result()
	
	if err == redis.Nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ShortID dosen't exists",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "problem connecting to redis server",
		})
		return
	}

	err = database.Client.Del(database.Ctx, shortID).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to Delete shortend link",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Shortened URL deleted successfully",
	})
}
