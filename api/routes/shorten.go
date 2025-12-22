package routes

import (
	"net/http"
	"os"
	"strconv"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Shrey7781/url-shortner/api/database"
	"github.com/Shrey7781/url-shortner/api/models"
	"github.com/Shrey7781/url-shortner/api/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func ShortenURL(c *gin.Context) {
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Parse JSON"})
		return
	}
	if body.URL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "URL field is required"})
        return
    }

	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()

	if err == redis.Nil || val == "" {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
			return
		}
	}
	body.URL = utils.EnsureHttpPrefix(body.URL)
	fmt.Printf("Validating URL: [%s]\n", body.URL)

	if !govalidator.IsRequestURL(body.URL) && !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL",
			"debug_received": body.URL,
		})
		return
	}

	if utils.IsServiceDomain(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "You can't hack this System :)",
		})
		return
	}

	

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	
	val, _ = database.Client.Get(database.Ctx, id).Result()

	if val != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "URL Custom Short Already Exists",
		})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}
	err = database.Client.Set(database.Ctx, id, body.URL, time.Duration(body.Expiry)*time.Hour).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp := models.Response{
		Expiry:          body.Expiry,
		XRateLimitReset: 30,
		XRateRemaining:  10,
		URL:             body.URL,
		CustomShort:     "",
	}
	r2.Decr(database.Ctx, c.ClientIP())
	val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
	resp.XRateLimitReset = int(ttl.Minutes())

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id
	c.JSON(http.StatusOK, resp)
}
