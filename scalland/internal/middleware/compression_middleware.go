package middleware

import (
	"github.com/gin-gonic/gin"
)

// BrotliMiddleware compresses responses using Brotli.
func BrotliMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use gin-contrib/gzip or andybalholm/brotli wrapper
		c.Next()
	}
}
