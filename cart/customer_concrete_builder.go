package cart

import "fmt"

// RegularCustomerConcreteBuilder is a concrete builder implementation for regular customers.
type RegularCustomerConcreteBuilder struct {
	cart *Cart
}

// NewRegularCustomerConcreteBuilder creates a new instance of RegularCustomerConcreteBuilder.
func NewRegularCustomerConcreteBuilder() *RegularCustomerConcreteBuilder {
	return &RegularCustomerConcreteBuilder{
		cart: &Cart{},
	}
}

// IntializeCart initializes the cart with default values.
func (c *RegularCustomerConcreteBuilder) IntializeCart(){
	c.cart.Items = make(map[string]Item)
}

// AddItem adds an item to the cart with the specified quantity.
func (c *RegularCustomerConcreteBuilder) AddItem(id, name, description string, price float64, quantity int) {
	item := Item{name, description, price, quantity}
	c.cart.Items[id] = item
}

// RemoveItem removes an item from the cart.
func (c *RegularCustomerConcreteBuilder) RemoveItem(item string) {
	if value, ok := c.cart.Items[item]; ok {
		value.Quantity -= 1
		if value.Quantity == 0 {
			delete(c.cart.Items, item)
		} else {
			c.cart.Items[item] = value
		}
	} else {
		fmt.Println("Item not found")
	}
}

// SetQuantity sets the quantity of an item in the cart.
func (c *RegularCustomerConcreteBuilder) SetQuantity(item string, quantity int) {
	if value, ok := c.cart.Items[item]; ok {
		value.Quantity = quantity
		if value.Quantity <= 0 {
			delete(c.cart.Items, item)
		} else {
			c.cart.Items[item] = value
		}
	} else {
		fmt.Println("Item not found")
	}
}

// ApplyDiscount applies a discount to the cart using the specified code.
func (c *RegularCustomerConcreteBuilder) ApplyDiscount() {
	c.cart.Discount = 10
	c.cart.TotalPrice = c.cart.TotalPrice - (c.cart.TotalPrice * c.cart.Discount / 100)
}

// SetUser sets the user ID for the cart.
func (c *RegularCustomerConcreteBuilder) SetUser(userID string) {
	c.cart.User = userID
}

// SetTax sets the tax for the cart.
func (c *RegularCustomerConcreteBuilder) SetTax() {
	c.cart.Taxes = 0.18*c.cart.TotalPrice + 0.05*c.cart.TotalPrice
}

// SetShippingDetails sets the shipping details for the cart.
func (c *RegularCustomerConcreteBuilder) SetShippingDetails(address, city, state, zip, country string) {
	c.cart.ShippingDetails = Shipping{address, city, state, zip, country}
}

// SetPaymentDetails sets the payment details for the cart.
func (c *RegularCustomerConcreteBuilder) SetPaymentDetails(cardNumber, expirationDate, cvv string) {
	c.cart.PaymentDetails = Payment{cardNumber, expirationDate, cvv}
}

// CalculateTotal calculates the total price of the cart.
func (c *RegularCustomerConcreteBuilder) CalculateTotal() float64 {
	for _, value := range c.cart.Items {
		c.cart.TotalPrice += value.Price * float64(value.Quantity)

	}
	return c.cart.TotalPrice
}

// Reset resets the cart to its initial state.
func (c *RegularCustomerConcreteBuilder) Reset() {
	c.cart = &Cart{}
}

// Build builds the cart and returns the final result.
func (c *RegularCustomerConcreteBuilder) Build() *Cart {
	return c.cart
}
