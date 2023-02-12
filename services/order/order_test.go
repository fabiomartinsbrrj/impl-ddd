package order

import (
	"testing"

	"github.com/fabiomartinsbrrj/tavern/domain/customer"
	"github.com/fabiomartinsbrrj/tavern/domain/product"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "healthy", 43.32)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := product.NewProduct("Peenuts", "Snackl", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
		beer, peenuts, wine,
	}

}

func TestOrder_NewOrderService(t *testing.T) {

	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
