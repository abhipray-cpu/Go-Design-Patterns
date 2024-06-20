package pet

// Director is responsible for constructing pets using a PetBuilder.
type Director struct {
	builder PetBuilder
}

// newDirector creates a new Director instance with the given PetBuilder.
func newDirector(p PetBuilder) *Director {
	return &Director{
		builder: p,
	}
}

// setBuilder sets the builder for the Director.
// The builder is responsible for constructing the Pet object.
func (d *Director) setBuilder(b PetBuilder) {
	d.builder = b
}

// buildAccessoryProduct builds an accessory product using the PetBuilder.
// It sets the attributes, category, brand, suitable for, and durability of the product.
// It returns the PetBuilder instance.
func (d *Director) buildAccessoryProduct(id,name string,price float64,description,vendor string,category string,brand string,suitable []string,durability string,weight,material,size,prodType string) PetBuilder {
	d.builder.SetAttributes(id,name,price,description,vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetSuitableFor(suitable)
	d.builder.SetDurability(durability)
	d.builder.SetWeight(weight)
	d.builder.SetMaterial(material)
	d.builder.SetSize(size)
	d.builder.SetType(prodType)
	return d.builder
}

// buildGroomingProduct builds a grooming product using the configured builder.
// It sets the attributes, category, brand, suitable for, volume, size, and type of the product.
// It returns the configured builder.
func (d *Director) buildGroomingProduct(id,name string,price float64,description,vendor string,category string,brand string,suitable []string,durability string,
	volume string,size string,prodType string, ingrdients []string) PetBuilder {
	d.builder.SetAttributes(id,name,price,description,vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetSuitableFor(suitable)
	d.builder.SetVolume(volume)
	d.builder.SetSize(size)
	d.builder.SetType(prodType)
	d.builder.SetIngredients(ingrdients)
	return d.builder
}

// buildFoodProduct builds a food product using the PetBuilder instance associated with the Director.
// It sets the attributes, category, brand, ingredients, and nutritional value of the product.
// It returns the PetBuilder instance after setting the attributes.
func (d *Director) buildFoodProduct(id,name string,price float64,description,vendor string,category string,brand string, ingredients []string, nutritionalValue string,weight,size,prodType string, suitable []string) PetBuilder {
	d.builder.SetAttributes(id,name,price,description,vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetIngredients(ingredients)
	d.builder.SetNutritionalValue(nutritionalValue)
	d.builder.SetWeight(weight)
	d.builder.SetSize(size)
	d.builder.SetType(prodType)
	d.builder.SetSuitableFor(suitable)
	return d.builder
}

// buildToyProduct builds a toy product using the current builder.
// It sets the attributes, category, brand, suitable for, and durability of the product.
// It returns the builder to allow further modifications or to retrieve the built product.
func (d *Director) buildToyProduct(id,name string,price float64,description,vendor string,category string,brand string, suitable []string, durability string,material,prodType,size string) PetBuilder {
	d.builder.SetAttributes(id,name,price,description,vendor)
	d.builder.SetCategory(category)
	d.builder.SetBrand(brand)
	d.builder.SetSuitableFor(suitable)
	d.builder.SetDurability(durability)
	d.builder.SetMaterial(material)
	d.builder.SetType(prodType)
	d.builder.SetSize(size)
	return d.builder
}