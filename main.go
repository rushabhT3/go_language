// main.go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	middlewares "my_gin_project/middlewares" // import your middleware package as middlewares
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Use the Logger middleware
    router.Use(middlewares.Logger()) // use the Logger function from the middlewares package

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, Gin with Middleware!",
        })
    })

    // Define a new route handler for the "/external" path
    router.GET("/external", func(c *gin.Context) {
        // Make a GET request to the external API
        resp, err := http.Get("https://crudcrud.com/api/d7b51435183d4369aeac85a3c2218642")
        if err != nil {
            // If there's an error, return a 500 status and an error message
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to make GET request",
            })
            return
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            // If there's an error, return a 500 status and an error message
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to read response body",
            })
            return
        }

        // Send the response body as the response
        c.String(http.StatusOK, string(body))
    })

    router.POST("/external", func(c *gin.Context) {
    // Generate a random number
    randomNum := rand.Intn(100) // generates a random integer between 0 and 100
    log.Println("Generated random number:", randomNum)

    // Create a map with the data you want to send
    data := map[string]string{
        "randomKey":   "key" + strconv.Itoa(randomNum), // generates a random key
        "randomValue": "value" + strconv.Itoa(randomNum), // generates a random value
    }
    log.Println("Created data map:", data)

    // Marshal the map to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Println("Failed to marshal data to JSON:", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to marshal data to JSON",
        })
        return
    }
    log.Println("Marshalled data to JSON:", string(jsonData))

    // Make a POST request to the external API
    resp, err := http.Post("https://crudcrud.com/api/d7b51435183d4369aeac85a3c2218642", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Println("Failed to make POST request:", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to make POST request",
        })
        return
    }
    log.Println("Made POST request, response status:", resp.Status)
    defer resp.Body.Close()

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Failed to read response body:", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to read response body",
        })
        return
    }
    log.Println("Read response body:", string(body))

    // Send the response body as the response
    c.String(http.StatusOK, string(body))
})

    router.Run(":8080")
}
