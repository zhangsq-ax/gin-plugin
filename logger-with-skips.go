package gin_plugin

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerWithSkips(skipPaths []string) gin.HandlerFunc {
	skipMap := make(map[string]struct{})
	for _, path := range skipPaths {
		skipMap[path] = struct{}{}
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		fullPath := c.FullPath()
		if _, ok := skipMap[fullPath]; ok {
			return
		}

		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.FullPath()

		log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s\n",
			start.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
