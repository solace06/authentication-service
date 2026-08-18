package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solace06/auth-service/pkg"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(10.0/60.0), 10)
var loginLimiter = pkg.NewRateLimiter()

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many signup requests",
			})
			return
		}

		c.Next()
	}
}

func LoginRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !loginLimiter.Allow(ip) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many request",
			})
			return
		}
		c.Next()
	}
}
