package product

import (
	"errors"
	"fvm/impl-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("no such product")
	ErrProductAlreadyExist = errors.New("there is already such an product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
