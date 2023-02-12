package memory

import (
	"sync"

	"github.com/fabiomartinsbrrj/tavern/domain/product"
	"github.com/google/uuid"
)

type memoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *memoryProductRepository {
	return &memoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mp *memoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product

	for _, product := range mp.products {
		products = append(products, product)
	}

	return products, nil
}

func (mp *memoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mp.products[id]; ok {
		return product, nil
	}

	return product.Product{}, product.ErrProductNotFound
}

func (mp *memoryProductRepository) Add(newprod product.Product) error {
	mp.Lock()
	defer mp.Unlock()

	if _, ok := mp.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mp.products[newprod.GetID()] = newprod

	return nil
}

func (mp *memoryProductRepository) Update(update product.Product) error {
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
