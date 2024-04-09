package models

import (
	"time"
)

// User defines the structure for a user model
type User struct {
    ID        string     `json:"id"`
    FirstName string   `json:"first_name"`
    LastName  string   `json:"last_name"`
    Email     string   `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}
