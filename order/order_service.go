package order

// OrderService defines the interface for managing orders.
type OrderService interface {
	// CreateOrder creates a new order with the given customer information, order items, and order type.
	// It returns the created order and an error if any.
	CreateOrder(customer CustomerInfo, items []OrderItem, orderType string) (*Order, error)

	// ProcessPayment processes the payment for the given order.
	// It returns an error if the payment processing fails.
	ProcessPayment(order *Order) error

	// ArrangeShipping arranges the shipping for the given order.
	// It returns an error if the shipping arrangement fails.
	ArrangeShipping(order *Order) error
}
