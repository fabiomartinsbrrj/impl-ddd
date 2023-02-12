package entity

import (
	"github.com/google/uuid"
)

//Package entities

// Person is an entity that represents a person in all domain
type Person struct {
	//ID an identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
