package health

// PersonalCareBuilder is a builder for creating personal care health products.
type PersonalCareBuilder struct {
	health Health
}

// newPersonalCareBuilder creates a new instance of PersonalCareBuilder.
func newPersonalCareBuilder() *PersonalCareBuilder {
	return &PersonalCareBuilder{}
}

// SetAttributes sets the attributes of the personal care health product.
func (p *PersonalCareBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	p.health.Id = id
	p.health.Name = name
	p.health.Price = price
	p.health.Description = description
	p.health.Vendor = vendor
}

// SetCategory sets the category of the personal care health product.
func (p *PersonalCareBuilder) SetCategory(category string) {
	p.health.Category = category
}

// SetIngredients sets the ingredients of the personal care health product.
func (p *PersonalCareBuilder) SetIngredients(ingredients []string) {
	p.health.Ingredients = ingredients
}

// SetBrand sets the brand of the personal care health product.
func (p *PersonalCareBuilder) SetBrand(brand string) {
	p.health.Brand = brand
}

// SetUsageInstructions sets the usage instructions of the personal care health product.
func (p *PersonalCareBuilder) SetUsageInstructions(usageInstructions []string) {
	p.health.UsageInstructions = usageInstructions
}

// SetType sets the type of the personal care health product.
func (p *PersonalCareBuilder) SetType(healthType string) {
	p.health.Type = healthType
}

// SetVolume sets the volume of the personal care health product.
func (p *PersonalCareBuilder) SetVolume(volume string) {
	p.health.Volume = volume
}

// SetForMenWomen sets whether the personal care health product is for men or women.
func (p *PersonalCareBuilder) SetForMenWomen(forMenWomen string) {
	p.health.ForMenWomen = forMenWomen
}

// GetItem returns the built personal care health product.
func (p *PersonalCareBuilder) GetItem() Health {
	return p.health
}

// SetSPF sets the SPF (Sun Protection Factor) of the personal care health product.
// This method is unused in the current implementation.
func (p *PersonalCareBuilder) SetSPF(spf string) {}

// SetFinish sets the finish of the personal care health product.
// This method is unused in the current implementation.
func (p *PersonalCareBuilder) SetFinish(finish string) {}

// SetSulfateFree sets whether the personal care health product is sulfate-free.
// This method is unused in the current implementation.
func (p *PersonalCareBuilder) SetSulfateFree(sulfateFree bool) {}

// SetSupplementForm sets the supplement form of the personal care health product.
// This method is unused in the current implementation.
func (p *PersonalCareBuilder) SetSupplementForm(supplementForm string) {}

// SetRecommendedDosage sets the recommended dosage of the personal care health product.
// This method is unused in the current implementation.
func (p *PersonalCareBuilder) SetRecommendedDosage(recommendedDosage string) {}

// SetColor sets the color of the personal care health product.
// This method is unused in the current implementation.
func (p *PersonalCareBuilder) SetColor(color string) {}
