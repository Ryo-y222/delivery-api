package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

// クロスオリジンリクエストを許可するミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowOrigin := os.Getenv("CORS_ORIGIN")
		if allowOrigin == "" {
			allowOrigin = "http://localhost:3000"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// プリフライトリクエストの処理
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
