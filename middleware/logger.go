package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger Middleware
func Logger() gin.HandlerFunc {
	fmt.Print("Milddleware")
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)
		log.Printf("Request %s %s - %v", c.Request.Method, c.Request.URL, latency)
	}
}
