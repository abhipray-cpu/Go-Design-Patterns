package cart

// CartDirector is a director that constructs a cart using a builder.
type Director struct {
	builder CartBuilder
}

// NewCartDirector creates a new CartDirector with the specified builder.
func NewCartDirector(builder CartBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// buildRegularCustomerCart builds a cart for a regular customer.
func (d *Director) buildCart(id, name, description string, price float64, quantity int) CartBuilder {
	d.builder.IntializeCart()
	d.builder.AddItem(id, name, description, price, quantity)
	d.builder.SetUser(id)
	d.builder.SetTax()
	d.builder.CalculateTotal()
	d.builder.ApplyDiscount()
	return d.builder
}
