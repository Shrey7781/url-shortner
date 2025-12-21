package routes

import (
	"net/http"
	"time"

	"github.com/Shrey7781/url-shortner/api/database"
	"github.com/Shrey7781/url-shortner/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context){
	shortID:=c.Param("shortID")
	var body models.Requset

	if err:=c.ShouldBind(&body);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Parse JSON",
		})
	} 

	r:=database.CreateClient(0)
	defer r.Close()

	val,err:= r.Get(database.Ctx,shortID).Result()
	if err!=nil || val==""{
		c.JSON(http.StatusNotFound, gin.H{
			"error":"ShortID dosen't exists",
		})
		return
	}
	err=r.Set(database.Ctx,shortID,body.URL,body.Expiry*3600*time.Second).Err()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Unable to update the shortend content",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "The Content has been Updated !!!",
	})
	 

}