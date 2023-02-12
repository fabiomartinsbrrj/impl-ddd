package main

import (
	"context"

	"github.com/fabiomartinsbrrj/tavern/domain/product"
	"github.com/fabiomartinsbrrj/tavern/services/order"
	"github.com/fabiomartinsbrrj/tavern/services/tavern"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	products := productsInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(ctx, "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern()
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Martins")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productsInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "healthy", 43.32)
	if err != nil {
		panic(err)
	}

	peenuts, err := product.NewProduct("Peenuts", "Snackl", 0.99)
	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		panic(err)
	}

	return []product.Product{
		beer, peenuts, wine,
	}

}
