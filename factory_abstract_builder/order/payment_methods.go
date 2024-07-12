package order

import "errors"

// PaymentMethod defines the interface for payment methods.
type PaymentMethod interface {
	ProcessPayment(details PaymentDetails) error
	CancelPayment(details PaymentDetails) error
}

// PaymentDetails encapsulates the payment information for an order.
type PaymentDetails struct {
	Method   string  // Payment method used, e.g., "Credit Card", "PayPal"
	Status   string  // Current status of the payment (e.g., "Pending", "Completed")
	Amount   float64 // Total amount paid
	Currency string  // Currency of the payment, e.g., "USD", "EUR"
}

// ProcessPayment processes the payment for an order.
// It updates the payment details with the provided method, amount, and currency.
// Returns the updated payment details and an error if the payment amount is invalid or the currency is empty.
func (p *PaymentDetails) ProcessPayment(method string, amount float64, currency string) (PaymentDetails, error) {
	if amount <= 0 {
		return PaymentDetails{}, errors.New("Invalid payment amount")
	}
	if currency == "" {
		return PaymentDetails{}, errors.New("Currency is required")
	}
	p.Method = method
	p.Status = "Completed"
	p.Amount = amount
	p.Currency = currency
	return *p, nil
}
