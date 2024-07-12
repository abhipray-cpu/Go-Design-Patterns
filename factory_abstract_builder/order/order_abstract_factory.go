package order

import "fmt"

// OrderFactory is an interface that defines the contract for creating orders.
type OrderFactory interface {
	CreateOrder(items []OrderItem, paymentMethod string, currency, address string, name, email string) (Order, error)
}

// getOrderFactory returns an instance of OrderFactory based on the orderType parameter.
func getOrderFactory(orderType string) (OrderFactory, error) {
	if orderType == "regular" {
		return &RegularOrderFactory{}, nil
	}
	if orderType == "gift" {
		return &PremiumOrderFactory{}, nil
	}
	return nil, fmt.Errorf("Wrong order type passed")
}
