package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Middleware logic before request
        c.Next()
        // Middleware logic after request
    }
}

func main() {
    router := gin.Default()
    // Use the Logger middleware
    router.Use(Logger())
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, Gin with Middleware!",
        })
    })
    router.Run(":8080")
}
