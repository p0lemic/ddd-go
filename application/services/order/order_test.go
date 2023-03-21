package order

import (
	"testing"

	"github.com/google/uuid"
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
func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)
	// Add Customer
	cust, err := customer.NewCustomer("Pepe")
	if err != nil {
		t.Error(err)
	}

	os, err := NewOrderService(
		WithMemoryCustomerRepository(cust),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Perform Order for one beer
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}

}
