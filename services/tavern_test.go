package services

import (
	"context"
	"fvm/impl-ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	/*os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)*/
	os, err := NewOrderService(
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("martins")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}
}
