package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Shrey7781/url-shortner/api/database"
	"github.com/go-redis/redis/v8"
)

func ResolveURL(c *gin.Context){
	shortID:= c.Param("shortID")
	val,err:= database.Client.Get(database.Ctx,shortID).Result()
	if err==redis.Nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found in database"})
        return
	}
	rInr := database.CreateClient(1)
	defer rInr.Close()
    _ = rInr.Incr(database.Ctx, "counter")
	c.Redirect(http.StatusMovedPermanently, val)
}