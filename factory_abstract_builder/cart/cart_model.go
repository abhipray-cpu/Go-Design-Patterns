package cart

import "time"

// Cart represents a shopping cart.
type Cart struct {
	Items           map[string]Item // Items stores the list of items in the cart.
	User            string          // User stores the username associated with the cart.
	TotalPrice      float64         // TotalPrice stores the total price of all items in the cart.
	Discount        float64         // Discount stores the discount applied to the cart.
	Taxes           float64         // Taxes stores the taxes applied to the cart.
	CreatedAt       time.Time       // CreatedAt stores the timestamp when the cart was created.
	ShippingDetails Shipping        // ShippingDetails stores the shipping information for the cart.
	PaymentDetails  Payment         // PaymentDetails stores the payment information for the cart.
}

type Item struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

type Shipping struct {
	Address string
	City    string
	State   string
	ZipCode string
	Country string
}

type Payment struct {
	CardNumber     string
	ExpirationDate string
	CVV            string
}
