package pet

// ToyBuilder is a builder for creating toy products.
type ToyBuilder struct {
	pets Pet
}

// newToyBuilder creates a new instance of ToyBuilder.
func newToyBuilder() *ToyBuilder {
	return &ToyBuilder{}
}

// SetAttributes sets the attributes of the toy product.
func (t *ToyBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	t.pets.ID = id
	t.pets.Name = name
	t.pets.Price = price
	t.pets.Description = description
	t.pets.Vendor = vendor
}

// SetCategory sets the category of the toy product.
func (t *ToyBuilder) SetCategory(category string) {
	t.pets.Category = category
}

// SetBrand sets the brand of the toy product.
func (t *ToyBuilder) SetBrand(brand string) {
	t.pets.Brand = brand
}

// SetWeight sets the weight of the toy product.
func (t *ToyBuilder) SetWeight(weight string) {
	t.pets.Weight = weight
}

// SetSuitableFor sets the suitable pets for the toy product.
func (t *ToyBuilder) SetSuitableFor(pets []string) {
	t.pets.SuitableFor = pets
}

// SetMaterial sets the material of the toy product.
func (t *ToyBuilder) SetMaterial(material string) {
	t.pets.Material = material
}

// SetDurability sets the durability of the toy product.
func (t *ToyBuilder) SetDurability(durability string) {
	t.pets.Durability = durability
}

// SetType sets the type of the toy product.
func (t *ToyBuilder) SetType(prodType string) {
	t.pets.Type = prodType
}

// SetSize sets the size of the toy product.
func (t *ToyBuilder) SetSize(size string) {
	t.pets.Size = size
}

// GetItem returns the built toy product.
func (t *ToyBuilder) GetItem() Pet {
	return t.pets
}

// unused methods from abstract interface

// SetIngredients sets the ingredients of the toy product.
func (t *ToyBuilder) SetIngredients(ingredients []string) {}

// SetNutritionalValue sets the nutritional value of the toy product.
func (t *ToyBuilder) SetNutritionalValue(value string) {}

// SetVolume sets the volume of the toy product.
func (t *ToyBuilder) SetVolume(volume string) {}
