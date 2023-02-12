package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("no such product")
	ErrProductAlreadyExist = errors.New("there is already such an product")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
