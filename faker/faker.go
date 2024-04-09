package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"my_gin_project/models"
)

func GenerateFakeUser() models.User {
	return models.User{
		ID:        uuid.New().String(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		CreatedAt: time.Now(),
	}
}

func GenerateFakeUsers(count int) []models.User {
  // creates a slice of User structs with length equal to count. make: initialize slices, maps, and channels.
	users := make([]models.User, count)
	for i := 0; i < count; i++ {
		users[i] = GenerateFakeUser()
	}
	return users
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	users := GenerateFakeUsers(10)
	for i, user := range users {
		fmt.Println("Fake User:", i+1)
		fmt.Printf("  ID:         %s\n", user.ID)
		fmt.Printf("  First Name: %s\n", user.FirstName)
		fmt.Printf("  Last Name:  %s\n", user.LastName)
		fmt.Printf("  Email:      %s\n", user.Email)
		fmt.Printf("  Created At: %s\n", user.CreatedAt.Format(time.RFC3339))

		_, err = db.Exec(`INSERT INTO users (id, first_name, last_name, email, created_at) VALUES ($1, $2, $3, $4, $5)`,
			user.ID, user.FirstName, user.LastName, user.Email, user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
