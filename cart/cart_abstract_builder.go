package cart

// CartBuilder is an interface that defines the methods for building a cart.
type CartBuilder interface {
	// IntializeCart initializes the cart with default values.
	IntializeCart()
	
	// AddItem adds an item to the cart with the specified quantity.
	AddItem(id, name, description string, price float64, quantity int)

	// RemoveItem removes an item from the cart.
	RemoveItem(item string)

	// SetQuantity sets the quantity of an item in the cart.
	SetQuantity(item string, quantity int)

	// ApplyDiscount applies a discount to the cart using the specified code.
	ApplyDiscount()

	// SetUser sets the user ID for the cart.
	SetUser(userID string)

	// SetCurrency sets the currency for the cart.
	SetTax()

	// setShippingDetails sets the shipping details for the cart.
	SetShippingDetails(address, city, state, zip, country string)

	// SetPaymentDetails sets the payment details for the cart.
	SetPaymentDetails(cardNumber, expirationDate, cvv string)

	// CalculateTotal calculates the total price of the items in the cart.
	CalculateTotal() float64

	// Reset resets the cart to its initial state.
	Reset()

	// Build builds and returns the cart.
	Build() *Cart
}

// getBuilder returns a CartBuilder based on the given builderType.
// It takes a string parameter builderType and returns a CartBuilder interface.
// The builderType can be one of the following values: "regular", "premium", or "guest".
// If the builderType is not recognized, it returns nil.
func getBuilder(builderType string) CartBuilder {
	switch builderType {
	case "regular":
		return NewRegularCustomerConcreteBuilder()
	case "premium":
		return NewPremiumCustomerConcreteBuilder()
	case "guest":
		return NewUnregisteredCustomerConcreteBuilder()
	default:
		return nil
	}
}
