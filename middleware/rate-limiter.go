package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	clients   = make(map[string]*client)
	mu        sync.Mutex
	rateLimit = rate.Every(1 * time.Second) // 1 request per second
	burst     = 5
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getClient(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Try again later.",
			})
			return
		}

		c.Next()
	}
}

func getClient(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if c, exists := clients[ip]; exists {
		c.lastSeen = time.Now()
		return c.limiter
	}

	limiter := rate.NewLimiter(rateLimit, burst)
	clients[ip] = &client{limiter: limiter, lastSeen: time.Now()}
	return limiter
}

func cleanupClients() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, c := range clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func init() {
	go cleanupClients()
}
