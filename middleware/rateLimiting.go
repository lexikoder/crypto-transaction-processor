package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware(rps float64, burst int) gin.HandlerFunc {
	// Store limiters per IP
	limiters := make(map[string]*rate.Limiter)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter, exists := limiters[ip]
		if !exists {
			// Create a new limiter: rps requests per second, burst capacity
			limiter = rate.NewLimiter(rate.Limit(rps), burst)
			limiters[ip] = limiter
		}

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
