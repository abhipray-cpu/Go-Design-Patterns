package health

// MakeupBuilder is a builder for creating Makeup products.
type MakeupBuilder struct {
	health Health
}

// newMakeUpBuilder creates a new instance of MakeupBuilder.
func newMakeUpBuilder() *MakeupBuilder {
	return &MakeupBuilder{}
}

// SetAttributes sets the attributes of the Makeup product.
func (m *MakeupBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	m.health.Id = id
	m.health.Name = name
	m.health.Price = price
	m.health.Description = description
	m.health.Vendor = vendor
}

// SetCategory sets the category of the Makeup product.
func (m *MakeupBuilder) SetCategory(category string) {
	m.health.Category = category
}

// SetIngredients sets the ingredients of the Makeup product.
func (m *MakeupBuilder) SetIngredients(ingredients []string) {
	m.health.Ingredients = ingredients
}

// SetBrand sets the brand of the Makeup product.
func (m *MakeupBuilder) SetBrand(brand string) {
	m.health.Brand = brand
}

// SetType sets the type of the Makeup product.
func (m *MakeupBuilder) SetType(healthType string) {
	m.health.Type = healthType
}

// SetVolume sets the volume of the Makeup product.
func (m *MakeupBuilder) SetVolume(volume string) {
	m.health.Volume = volume
}

// SetColor sets the color of the Makeup product.
func (m *MakeupBuilder) SetColor(color string) {
	m.health.Color = color
}

// SetFinish sets the finish of the Makeup product.
func (m *MakeupBuilder) SetFinish(finish string) {
	m.health.Finish = finish
}

// GetItem returns the built Makeup product.
func (m *MakeupBuilder) GetItem() Health {
	return m.health
}

// unused methods from abstract interface

// SetSPF sets the SPF of the Makeup product.
func (m *MakeupBuilder) SetSPF(spf string) {}

// SetUsageInstructions sets the usage instructions of the Makeup product.
func (m *MakeupBuilder) SetUsageInstructions(usageInstructions []string) {}

// SetForMenWomen sets whether the Makeup product is for men or women.
func (m *MakeupBuilder) SetForMenWomen(forMenWomen string) {}

// SetSulfateFree sets whether the Makeup product is sulfate-free.
func (m *MakeupBuilder) SetSulfateFree(sulfateFree bool) {}

// SetRecommendedDosage sets the recommended dosage of the Makeup product.
func (m *MakeupBuilder) SetRecommendedDosage(recommendedDosage string) {}

// SetSupplementForm sets the supplement form of the Makeup product.
func (m *MakeupBuilder) SetSupplementForm(supplementForm string) {}
