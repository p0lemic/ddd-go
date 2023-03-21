package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p0lemic/ddd-go/domain"
)

var (
	ErrInvalidPerson = errors.New("a customer needs a valid person")
)

type Customer struct {
	person       *domain.Person
	products     []*domain.Item
	transactions []domain.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &domain.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*domain.Item, 0),
		transactions: make([]domain.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &domain.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &domain.Person{}
	}
	c.person.Name = name
}

// SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}
