package cart

// createNormalCustomerCart creates a cart for a normal customer using the regular builder.
// It returns the created cart.
func createCart(userType string, id, name, description string, price float64, quantity int) CartBuilder {
	var builder CartBuilder
	switch userType {
	case "regular":
		builder = getBuilder("regular")
	case "premium":
		builder = getBuilder("premium")
	case "guest":
		builder = getBuilder("guest")
	default:
		return nil
	}

	director := NewCartDirector(builder)
	cart := director.buildCart(id, name, description, price, quantity)
	return cart
}
