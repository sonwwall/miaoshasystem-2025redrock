package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RateLimitMiddleware(limit int) gin.HandlerFunc { //反脚本哥中间件
	requestMap := make(map[string]int)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if requestMap[ip] >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "😠检测到你滴脚本辽！重拳出击👊💥 一个都莫跑🏃‍♂️🚫"})
			c.Abort()
			return
		}
		requestMap[ip] = requestMap[ip] + 1
		go func() {
			time.Sleep(1 * time.Second)
			delete(requestMap, ip)
		}()

		c.Next()
	}
}
