package services

import (
	"context"
	"fvm/impl-ddd/aggregate"
	"fvm/impl-ddd/domain/customer"
	"fvm/impl-ddd/domain/customer/memory"
	"fvm/impl-ddd/domain/customer/mongo"
	"fvm/impl-ddd/domain/product"
	prodmem "fvm/impl-ddd/domain/product/memory"
	"log"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {

	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

/* Exemplo
NewOrderService(
	WithMongoCustomerRepository(),
	WithMemoryProductRepository(),
	WithLogging("debug")
)
*/

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}

		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	//Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	//Get each Product
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has order %d product", c.GetID(), len(products))

	return total, nil
}
