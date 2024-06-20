package toys

type Director struct {
	builder ToyBuilder
}

// newDirector creates a new Director instance with the given ToyBuilder.
func newDirector(b ToyBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// setBuilder sets the ToyBuilder for the Director.
// The ToyBuilder is responsible for constructing the toy.
func (d *Director) setBuilder(b ToyBuilder) {
	d.builder = b
}

// buildVideoToy builds a board game using the configured ToyBuilder.
// It sets the attributes, category, age range, brand, weight, and vendor of the toy.
// It returns the configured ToyBuilder.
func (d *Director) buildBoardToy(id, name string, price float64, description string, category string, minAge, maxAge int, brand string, weight string, vendor string, prodType, dimension string, suitable []string) ToyBuilder {
	d.builder.SetAttributes(id, name, price, description)
	d.builder.SetCategory(category)
	d.builder.SetAgeRange(minAge, maxAge)
	d.builder.SetBrand(brand)
	d.builder.SetWeight(weight)
	d.builder.SetVendor(vendor)
	d.builder.SetType(prodType)
	d.builder.SetDimensions(dimension)
	d.builder.SetSuitableFor(suitable)
	return d.builder

}

// buildVideoToy builds a video toy using the configured ToyBuilder.
// It sets the attributes, category, age range, brand, weight, and vendor of the toy.
// It returns the configured ToyBuilder.
func (d *Director) buildVideoToy(id, name string, price float64, description string, category string, minAge, maxAge int, brand string, weight string, vendor string, suitable []string) ToyBuilder {
	d.builder.SetAttributes(id, name, price, description)
	d.builder.SetCategory(category)
	d.builder.SetAgeRange(minAge, maxAge)
	d.builder.SetBrand(brand)
	d.builder.SetWeight(weight)
	d.builder.SetSuitableFor(suitable)
	d.builder.SetVendor(vendor)
	return d.builder
}

// builActionToy builds an action toy using the configured ToyBuilder.
// It sets the attributes, category, age range, brand, weight, material, color, and vendor of the toy.
// It returns the configured ToyBuilder.
func (d *Director) buildActionToy(id, name string, price float64, description string, category string, minAge, maxAge int, brand string, weight string, vendor string, material, color string, size, prodType, dimension string, suitable []string) ToyBuilder {
	d.builder.SetAttributes(id, name, price, description)
	d.builder.SetCategory(category)
	d.builder.SetAgeRange(minAge, maxAge)
	d.builder.SetBrand(brand)
	d.builder.SetWeight(weight)
	d.builder.SetMaterial(material)
	d.builder.SetColor(color)
	d.builder.SetSize(size)
	d.builder.SetType(prodType)
	d.builder.SetDimensions(dimension)
	d.builder.SetSuitableFor(suitable)
	d.builder.SetVendor(vendor)
	return d.builder
}
