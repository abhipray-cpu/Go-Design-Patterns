package pet

// AccessoryBuilder is a builder for creating pet accessories.
type AccessoryBuilder struct {
	pets Pet
}

// newAccessoryBuilder creates a new instance of AccessoryBuilder.
func newAccessoryBuilder() *AccessoryBuilder {
	return &AccessoryBuilder{}
}

// SetAttributes sets the attributes of the pet accessory.
func (a *AccessoryBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	a.pets.ID = id
	a.pets.Name = name
	a.pets.Price = price
	a.pets.Description = description
	a.pets.Vendor = vendor
}

// SetCategory sets the category of the pet accessory.
func (a *AccessoryBuilder) SetCategory(category string) {
	a.pets.Category = category
}

// SetBrand sets the brand of the pet accessory.
func (a *AccessoryBuilder) SetBrand(brand string) {
	a.pets.Brand = brand
}

// SetSuitableFor sets the suitable for field of the pet accessory.
func (a *AccessoryBuilder) SetSuitableFor(suitableFor []string) {
	a.pets.SuitableFor = suitableFor
}

// SetDurability sets the durability of the pet accessory.
func (a *AccessoryBuilder) SetDurability(durability string) {
	a.pets.Durability = durability
}

// SetWeight sets the weight of the pet accessory.
func (a *AccessoryBuilder) SetWeight(weight string) {
	a.pets.Weight = weight
}

// SetMaterial sets the material of the pet accessory.
func (a *AccessoryBuilder) SetMaterial(material string) {
	a.pets.Material = material
}

// SetSize sets the size of the pet accessory.
func (a *AccessoryBuilder) SetSize(size string) {
	a.pets.Size = size
}

// SetType sets the type of the pet accessory.
func (a *AccessoryBuilder) SetType(prodType string) {
	a.pets.Type = prodType
}

// GetItem returns the built pet accessory.
func (a *AccessoryBuilder) GetItem() Pet {
	return a.pets
}

// unused methods from abstract interface

// SetIngredients sets the ingredients of the pet accessory.
func (a *AccessoryBuilder) SetIngredients(ingredients []string) {}

// SetVolume sets the volume of the pet accessory.
func (a *AccessoryBuilder) SetVolume(volume string) {}

// SetNutritionalValue sets the nutritional value of the pet accessory.
func (a *AccessoryBuilder) SetNutritionalValue(value string) {}
