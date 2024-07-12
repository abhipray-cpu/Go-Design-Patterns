package order

import (
	"errors"
	"fmt"
	"time"
)

// PremiumOrderFactory is a factory for creating premium orders.
type PremiumOrderFactory struct{}

// CreateOrder creates a new order with the given items, payment method, currency, address, name, and email.
// It returns the created order and an error if any.
func (r *PremiumOrderFactory) CreateOrder(items []OrderItem, paymentMethod string, currency, address string, name, email string) (Order, error) {
	item_count := len(items)
	if item_count == 0 {
		return Order{}, errors.New("Items are required")
	}
	if address == "" {
		return Order{}, errors.New("Address is required")
	}
	total_amount := 0.0
	for item := range items {
		total_amount += items[item].Price * float64(items[item].Quantity)
	}

	shipping := ShippingDetails{}
	shipping_details, err := shipping.Ship("Express", address, items)

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
