package application

import "github.com/google/uuid"

type User struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Biography string    `json:"biography"`
	Id        uuid.UUID `json:"id"`
}