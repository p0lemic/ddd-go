package tavern

import (
	"testing"

	"github.com/google/uuid"
	"github.com/p0lemic/ddd-go/application/services/order"
	"github.com/p0lemic/ddd-go/domain/customer"
	"github.com/p0lemic/ddd-go/domain/product"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)
	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(cust),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
