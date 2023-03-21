package order

import (
	"log"

	"github.com/google/uuid"
	"github.com/p0lemic/ddd-go/domain/customer"
	"github.com/p0lemic/ddd-go/domain/product"
	"github.com/p0lemic/ddd-go/infrastructure/persistence/customer/memory"
	productmemory "github.com/p0lemic/ddd-go/infrastructure/persistence/product/memory"
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

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository(customer customer.Customer) OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	return func(os *OrderService) error {
		cr := memory.New()
		cr.Add(customer)

		os.customers = cr
		return nil
	}
}

// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := productmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

// CreateOrder will chaintogether all repositories to create a order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product, Ouchie, We need a ProductRepository
	var products []product.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	customer, err := customer.NewCustomer(name)

	if err != nil {
		return uuid.UUID{}, err
	}

	o.customers.Add(customer)

	return customer.GetID(), nil
}
