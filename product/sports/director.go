package sports

// Director is responsible for constructing a product using the builder pattern.
type Director struct {
	builder SportsBuilder
}

// newDirector creates a new Director instance with the given SportsBuilder.
func newDirector(s SportsBuilder) *Director {
	return &Director{
		builder: s,
	}
}

// setBuilder sets the builder for the Director.
// It takes a SportsBuilder as a parameter and assigns it to the Director's builder field.
func (d *Director) setBuilder(b SportsBuilder) {
	d.builder = b
}

// buildSportsProduct builds a sports product using the configured builder.
// It sets the attributes, category, age group, dimensions, durability,
// material, weight, and vendor of the product.
// It returns the configured builder.
// buildSportsProduct builds a sports product using the provided attributes and returns a SportsBuilder.
func (d *Director) buildSportsProduct(id,name string,price float64,description string,category string,minAge,maxAge int,dimension string,durability string,material string,weight string,vendor string,prodType,size,brand,color string,suitable []string) SportsBuilder {
	d.builder.SetAttributes(id,name,price,description)
	d.builder.SetCategory(category)
	d.builder.SetAgeGroup(minAge,maxAge)
	d.builder.SetDimensions(dimension)
	d.builder.SetDurability(durability)
	d.builder.SetMaterial(material)
	d.builder.SetWeight(weight)
	d.builder.SetType(prodType)
	d.builder.SetSize(size)
	d.builder.SetBrand(brand)
	d.builder.SetColor(color)
	d.builder.SetSuitableFor(suitable)
	d.builder.SetVendor(vendor)
	return d.builder

}
