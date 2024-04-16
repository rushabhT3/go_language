package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "my_gin_project/db"
	"my_gin_project/models"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

// ExternalHandler handles the "/external" route
func ExternalHandler(c *gin.Context) {
	// Connect to Redis
	rdb, err := db.ConnectRedis()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Redis"})
		return
	}

	// Try to get the data from Redis first
	val, err := rdb.Get(db.Ctx, "users").Result()
	if err == redis.Nil {
		// Data not found in Redis, fetch from database
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

		// Convert the users slice to JSON
		usersJson, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert users to JSON"})
			return
		}

		// Store the JSON string in Redis
		err = rdb.Set(db.Ctx, "users", usersJson, 0).Err()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store data in Redis"})
			return
		}

		// Return the users as JSON
		c.JSON(http.StatusOK, users)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data from Redis"})
	} else {
		// Data found in Redis, return it
		c.JSON(http.StatusOK, val)
	}
}
