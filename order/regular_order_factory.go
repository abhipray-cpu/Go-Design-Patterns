package order

import (
	"errors"
	"fmt"
	"time"
)

// RegularOrderFactory is a factory for creating regular orders.
type RegularOrderFactory struct{}

// CreateOrder creates a regular order with the given parameters.
// It returns the created order and an error if any.
func (r *RegularOrderFactory) CreateOrder(items []OrderItem, paymentMethod string, currency, address string, name, email string) (Order, error) {
	item_count := len(items)
	if item_count == 0 {
		return Order{}, errors.New("items are required")
	}
	if address == "" {
		return Order{}, errors.New("address is required")
	}
	total_amount := 0.0
	for item := range items {
		total_amount += items[item].Price * float64(items[item].Quantity)
	}

	shipping := ShippingDetails{}
	shipping_details, err := shipping.Ship("Standard", address, items)

	if err != nil {
		return Order{}, err
	}
	payment := PaymentDetails{}
	payment_details, err := payment.ProcessPayment(paymentMethod, total_amount, currency)

	if err != nil {
		return Order{}, err
	}
	currentTime := time.Now()
	uniqueString := fmt.Sprintf("%d", currentTime.Unix())
	order := Order{
		uniqueString,
		CustomerInfo{name, email, address},
		items,
		payment_details,
		shipping_details,
		currentTime,
		shipping_details.ExpectedDelivery,
		"Processing",
	}

	return order, nil
}
