package cart

import "fmt"

// UnregisteredCustomerCartBuilder is a concrete builder for creating a cart for unregistered customers.
type UnregisteredCustomerCartBuilder struct {
	cart *Cart
}

// NewUnregisteredCustomerConcreteBuilder creates a new instance of UnregisteredCustomerCartBuilder.
func NewUnregisteredCustomerConcreteBuilder() *UnregisteredCustomerCartBuilder {
	return &UnregisteredCustomerCartBuilder{
		cart: &Cart{},
	}
}

// IntializeCart initializes the cart with default values.
func (c *UnregisteredCustomerCartBuilder) IntializeCart(){
	c.cart.Items = make(map[string]Item)
}

// AddItem adds an item to the cart with the specified quantity.
func (c *UnregisteredCustomerCartBuilder) AddItem(id, name, description string, price float64, quantity int) {
	item := Item{name, description, price, quantity}
	c.cart.Items[id] = item
}

// RemoveItem removes an item from the cart.
func (c *UnregisteredCustomerCartBuilder) RemoveItem(item string) {
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
func (c *UnregisteredCustomerCartBuilder) SetQuantity(item string, quantity int) {
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
func (c *UnregisteredCustomerCartBuilder) ApplyDiscount() {
	c.cart.Discount = 0
}

// SetUser sets the user ID for the cart.
func (c *UnregisteredCustomerCartBuilder) SetUser(userID string) {
	c.cart.User = userID
}

// SetTax sets the tax for the cart.
func (c *UnregisteredCustomerCartBuilder) SetTax() {
	c.cart.Taxes = 0.18*c.cart.TotalPrice + 0.1*c.cart.TotalPrice
}

// SetShippingDetails sets the shipping details for the cart.
func (c *UnregisteredCustomerCartBuilder) SetShippingDetails(address, city, state, zip, country string) {
	c.cart.ShippingDetails = Shipping{address, city, state, zip, country}
}

// SetPaymentDetails sets the payment details for the cart.
func (c *UnregisteredCustomerCartBuilder) SetPaymentDetails(cardNumber, expirationDate, cvv string) {
	c.cart.PaymentDetails = Payment{cardNumber, expirationDate, cvv}
}

// CalculateTotal calculates the total price of the cart.
func (c *UnregisteredCustomerCartBuilder) CalculateTotal() float64 {
	for _, value := range c.cart.Items {
		c.cart.TotalPrice += value.Price * float64(value.Quantity)

	}
	return c.cart.TotalPrice
}

// Reset resets the cart to its initial state.
func (c *UnregisteredCustomerCartBuilder) Reset() {
	c.cart = &Cart{}
}

// Build builds the cart and returns the final result.
func (c *UnregisteredCustomerCartBuilder) Build() *Cart {
	return c.cart
}
