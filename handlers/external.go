// handlers/external.go
package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExternalHandler(c *gin.Context) {
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
}
