package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3) // Rate limit of 1 request per second with a burst of 3 requests

func Ratelimited() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			// Exceeded request limit
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests, please wait and try again"})
			return
		}
		c.Next()
	}
}
