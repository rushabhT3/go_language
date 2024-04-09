package handlers

import (
	"net/http"

	db "my_gin_project/db"
	"my_gin_project/models"

	"github.com/gin-gonic/gin"
)

// ExternalHandler handles the "/external" route
func ExternalHandler(c *gin.Context) {
	// Connect to the database
	dbConn, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer dbConn.Close()

	// Query the database for all users
	rows, err := dbConn.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer rows.Close()

	// Scan the rows into a slice of User structs
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		users = append(users, user)
	}

	// Return the users as JSON
	c.JSON(http.StatusOK, users)
}
