package memory

import (
	"fvm/impl-ddd/aggregate"
	"fvm/impl-ddd/domain/product"
	"sync"

	"github.com/google/uuid"
)

type memoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *memoryProductRepository {
	return &memoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mp *memoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mp.products {
		products = append(products, product)
	}

	return products, nil
}

func (mp *memoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mp.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (mp *memoryProductRepository) Add(newprod aggregate.Product) error {
	mp.Lock()
	defer mp.Unlock()

	if _, ok := mp.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mp.products[newprod.GetID()] = newprod

	return nil
}

func (mp *memoryProductRepository) Update(update aggregate.Product) error {
	mp.Lock()
	defer mp.Unlock()

	if _, ok := mp.products[update.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mp.products[update.GetID()] = update

	return nil
}

func (mp *memoryProductRepository) Delete(id uuid.UUID) error {
	mp.Lock()
	defer mp.Unlock()

	if _, ok := mp.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mp.products, id)

	return nil
}
