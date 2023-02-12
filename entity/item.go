package entity

import "github.com/google/uuid"

type Item struct {
	//ID an identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
