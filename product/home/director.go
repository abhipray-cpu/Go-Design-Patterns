package home

// Director is a struct that defines the director of the builder pattern.
// It has a builder attribute that is of type HomeBuilder interface.
// The director is responsible for constructing the product using the builder.
type Director struct {
	builder HomeBuilder
}

// NewDirector is a function that returns a new instance of the Director struct.
// It takes a HomeBuilder instance as a parameter and returns a pointer to the Director struct.
func NewDirector(b HomeBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// setBuilder is a method of the Director struct that sets the builder attribute to the provided HomeBuilder instance.
// It takes a HomeBuilder instance as a parameter and does not return anything.
func (d *Director) setBuilder(b HomeBuilder) {
	d.builder = b
}

// buildFurniture builds a furniture product using the provided attributes and returns the HomeBuilder instance.
func (d *Director) buildFurniture(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimension string, assembly bool, capacity string) HomeBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetColor(color)
	d.builder.SetMaterial(material)
	d.builder.SetWeight(weight)
	d.builder.SetStyle(style)
	d.builder.SetDimensions(dimension)
	d.builder.SetAssemblyRequired(assembly)
	d.builder.SetCapacity(capacity)
	return d.builder
}

// buildKitchen builds a kitchen product using the provided attributes and returns the HomeBuilder instance.
func (d *Director) buildKitchen(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimension, capacity string, safe bool, assembly bool, power string) HomeBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetColor(color)
	d.builder.SetMaterial(material)
	d.builder.SetWeight(weight)
	d.builder.SetStyle(style)
	d.builder.SetDimensions(dimension)
	d.builder.SetCapacity(capacity)
	d.builder.SetDishWasherSafe(safe)
	d.builder.SetAssemblyRequired(assembly)
	d.builder.SetPowerSource(power)
	return d.builder
}

// buildDecor builds a decor product using the provided attributes and returns the HomeBuilder instance.
func (d *Director) buildDecor(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimension string, assembly bool) HomeBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetColor(color)
	d.builder.SetMaterial(material)
	d.builder.SetWeight(weight)
	d.builder.SetStyle(style)
	d.builder.SetDimensions(dimension)
	d.builder.SetAssemblyRequired(assembly)
	return d.builder
}

// buildBath builds a bath product using the provided attributes and returns the HomeBuilder instance.
func (d *Director) buildBath(id, name string, price float64, description, vendor string, category string, brand, prodType string, volume, scent, skinType string, material, dimension, weight, color, capacity, style string, assembly bool) HomeBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetVolume(volume)
	d.builder.SetScent(scent)
	d.builder.SetMaterial(material)
	d.builder.SetDimensions(dimension)
	d.builder.SetWeight(weight)
	d.builder.SetColor(color)
	d.builder.SetCapacity(capacity)
	d.builder.SetStyle(style)
	d.builder.SetAssemblyRequired(assembly)
	return d.builder
}

// buildTools builds a tools product using the provided attributes and returns the HomeBuilder instance.
func (d *Director) buildTools(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimensions, powerSource string) HomeBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetColor(color)
	d.builder.SetMaterial(material)
	d.builder.SetWeight(weight)
	d.builder.SetStyle(style)
	d.builder.SetDimensions(dimensions)
	d.builder.SetPowerSource(powerSource)
	return d.builder
}
