package order

import "time"

// CustomerInfo represents basic customer details.
type CustomerInfo struct {
	Name    string // Customer's full name
	Email   string // Customer's email address
	Address string // Primary address for the customer
}

// OrderItem represents an individual item within an order.
type OrderItem struct {
	ProductID   string  // Unique identifier for the product
	ProductName string  // Name of the product
	Quantity    int     // Number of units of the product
	Price       float64 // Price per unit of the product
}

// Order represents the entire order placed by a customer.
type Order struct {
	ID                string          // Unique identifier for the order
	Customer          CustomerInfo    // Customer information for the order
	Items             []OrderItem     // List of items included in the order
	Payment           PaymentDetails  // Payment details for the order
	Shipping          ShippingDetails // Shipping details for the order
	OrderDate         time.Time       // Date when the order was placed
	EstimatedDelivery time.Time       // Estimated delivery date for the order
	Status            string          // Current status of the order, e.g., "Processing", "Shipped"
}
