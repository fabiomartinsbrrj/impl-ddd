//Package memory is a in-memory implementation of Customer repository

package memory

import (
	"fmt"

	"sync"

	"github.com/fabiomartinsbrrj/tavern/domain/customer"
	"github.com/google/uuid"
)

type memoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *memoryRepository {
	return &memoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *memoryRepository) Get(id uuid.UUID) (customer.Customer, error) {

	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *memoryRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}
	//make sure customer is already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exist : %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *memoryRepository) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist : %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}
