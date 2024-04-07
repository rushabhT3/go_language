// middleware/logger.go
package middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Middleware logic before request
        c.Next()
        // Middleware logic after request
    }
}
