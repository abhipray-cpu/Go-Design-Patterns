package pet

// GroomingBuilder is a builder for creating grooming products.
type GroomingBuilder struct {
	pet Pet
}

// newGroomingBuilder creates a new instance of GroomingBuilder.
func newGroomingBuilder() *GroomingBuilder {
	return &GroomingBuilder{}
}

// SetAttributes sets the attributes of the grooming product.
func (g *GroomingBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	g.pet.ID = id
	g.pet.Name = name
	g.pet.Price = price
	g.pet.Description = description
	g.pet.Vendor = vendor
}

// SetCategory sets the category of the grooming product.
func (g *GroomingBuilder) SetCategory(category string) {
	g.pet.Category = category
}

// SetBrand sets the brand of the grooming product.
func (g *GroomingBuilder) SetBrand(brand string) {
	g.pet.Brand = brand
}

// SetSuitableFor sets the suitable for field of the grooming product.
func (g *GroomingBuilder) SetSuitableFor(suitableFor []string) {
	g.pet.SuitableFor = suitableFor
}

// SetVolume sets the volume of the grooming product.
func (g *GroomingBuilder) SetVolume(volume string) {
	g.pet.Volume = volume
}

// SetSize sets the size of the grooming product.
func (g *GroomingBuilder) SetSize(size string) {
	g.pet.Size = size
}

// SetType sets the type of the grooming product.
func (g *GroomingBuilder) SetType(prodType string) {
	g.pet.Type = prodType
}

// SetIngredients sets the ingredients of the grooming product.
func (g *GroomingBuilder) SetIngredients(ingredients []string) {
	g.pet.Ingredients = ingredients
}

// GetItem returns the built grooming product.
func (g *GroomingBuilder) GetItem() Pet {
	return g.pet
}

// unused methods from abstract interface

// SetNutritionalValue is an unused method from the abstract interface.
func (g *GroomingBuilder) SetNutritionalValue(value string) {}

// SetWeight is an unused method from the abstract interface.
func (g *GroomingBuilder) SetWeight(weight string) {}

// SetMaterial is an unused method from the abstract interface.
func (g *GroomingBuilder) SetMaterial(material string) {}

// SetDurability is an unused method from the abstract interface.
func (g *GroomingBuilder) SetDurability(durability string) {}
