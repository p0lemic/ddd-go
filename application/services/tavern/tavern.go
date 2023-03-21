package tavern

import (
	"log"

	"github.com/google/uuid"
	"github.com/p0lemic/ddd-go/application/services/order"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService   *order.OrderService
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the Customer: %0.0f", price)

	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	return nil
}
