package fashion

// Director is responsible for constructing fashion products using a FashionBuilder.
type Director struct {
	builder FashionBuilder
}

// NewDirector creates a new Director instance with the given FashionBuilder.
func NewDirector(builder FashionBuilder) *Director {
	return &Director{builder: builder}
}

// setBuilder sets the builder for the Director.
// The builder parameter should implement the FashionBuilder interface.
func (d *Director) setBuilder(builder FashionBuilder) {
	d.builder = builder
}

// buildClothing builds a fashion product using the provided attributes and returns a FashionBuilder.
// It sets the attributes, brand, category, material, size, color, gender, style, and product type of the fashion product.
func (d *Director) buildClothing(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string) FashionBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetMaterial(material)
	d.builder.SetSize(size)
	d.builder.SetColor(color)
	d.builder.SetGender(gender)
	d.builder.SetStyle(style)
	d.builder.SetType(prodType)
	return d.builder

}

// buildShoes builds a fashion product of type shoes using the provided attributes and returns a FashionBuilder.
// The function sets the attributes such as id, name, price, description, and vendor using the provided values.
// It also sets the brand, category, material, size, color, gender, style, and product type of the fashion product.
func (d *Director) buildShoes(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string) FashionBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetMaterial(material)
	d.builder.SetSize(size)
	d.builder.SetColor(color)
	d.builder.SetGender(gender)
	d.builder.SetStyle(style)
	d.builder.SetType(prodType)
	return d.builder
}

// buildJewelry builds a jewelry product using the provided attributes and returns a FashionBuilder.
// The function sets the attributes such as id, name, price, description, and vendor using the provided values.
// It also sets the brand, category, material size, color
func (d *Director) buildJewelry(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string, gemStone string) FashionBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetMaterial(material)
	d.builder.SetSize(size)
	d.builder.SetColor(color)
	d.builder.SetGender(gender)
	d.builder.SetStyle(style)
	d.builder.SetType(prodType)
	d.builder.SetGemStone(gemStone)
	return d.builder
}

// buildBag builds a fashion product of type bag using the provided attributes and returns a FashionBuilder.
// The function sets the attributes such as id, name, price, description, and vendor using the provided values.
// It also sets the brand, category, material size, color
func (d *Director) buildBag(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string) FashionBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetMaterial(material)
	d.builder.SetSize(size)
	d.builder.SetColor(color)
	d.builder.SetStyle(style)
	d.builder.SetType(prodType)
	return d.builder
}

// buildWatch builds a fashion product of type watch using the provided attributes and returns a FashionBuilder.
// The function sets the attributes such as id, name, price, description, and vendor using the provided values.
// It also sets the brand, category, material size, color
func (d *Director) buildWatch(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string, resistance string) FashionBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetMaterial(material)
	d.builder.SetSize(size)
	d.builder.SetColor(color)
	d.builder.SetGender(gender)
	d.builder.SetStyle(style)
	d.builder.SetType(prodType)
	d.builder.SetWaterResistance(resistance)
	return d.builder
}
