package pet

// FoodBuilder is a builder for creating Pet objects with specific attributes.
type FoodBuilder struct {
	pets Pet
}

// newFoodBuilder creates a new instance of FoodBuilder.
func newFoodBuilder() *FoodBuilder {
	return &FoodBuilder{}
}

// SetAttributes sets the attributes of the Pet object.
func (f *FoodBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	f.pets.ID = id
	f.pets.Name = name
	f.pets.Price = price
	f.pets.Description = description
	f.pets.Vendor = vendor
}

// SetCategory sets the category of the Pet object.
func (f *FoodBuilder) SetCategory(category string) {
	f.pets.Category = category
}

// SetBrand sets the brand of the Pet object.
func (f *FoodBuilder) SetBrand(brand string) {
	f.pets.Brand = brand
}

// SetIngredients sets the ingredients of the Pet object.
func (f *FoodBuilder) SetIngredients(ingredients []string) {
	f.pets.Ingredients = ingredients
}

// SetNutritionalValue sets the nutritional value of the Pet object.
func (f *FoodBuilder) SetNutritionalValue(value string) {
	f.pets.NutritionalValue = value
}

// SetWeight sets the weight of the Pet object.
func (f *FoodBuilder) SetWeight(weight string) {
	f.pets.Weight = weight
}

// SetSize sets the size of the Pet object.
func (f *FoodBuilder) SetSize(size string) {
	f.pets.Size = size
}

// SetSuitableFor sets the suitable pets for the Pet object.
func (f *FoodBuilder) SetSuitableFor(pets []string) {
	f.pets.SuitableFor = pets
}

// SetType sets the type of the Pet object.
func (f *FoodBuilder) SetType(prodType string) {
	f.pets.Type = prodType
}

// GetItem returns the Pet object.
func (f *FoodBuilder) GetItem() Pet {
	return f.pets
}

// unused methods from abstract interface

// SetVolume sets the volume of the Pet object.
func (f *FoodBuilder) SetVolume(volume string) {}

// SetMaterial sets the material of the Pet object.
func (f *FoodBuilder) SetMaterial(material string) {}

// SetDurability sets the durability of the Pet object.
func (f *FoodBuilder) SetDurability(durability string) {}
