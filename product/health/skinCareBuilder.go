package health

// SkinCareBuilder is a builder for creating SkinCare products.
type SkinCareBuilder struct {
	health Health
}

// newSkinCareBuilder creates a new instance of SkinCareBuilder.
func newSkinCareBuilder() *SkinCareBuilder {
	return &SkinCareBuilder{}
}

// SetAttributes sets the attributes of the SkinCare product.
func (s *SkinCareBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	s.health.Id = id
	s.health.Name = name
	s.health.Price = price
	s.health.Description = description
	s.health.Vendor = vendor
}

// SetCategory sets the category of the SkinCare product.
func (s *SkinCareBuilder) SetCategory(category string) {
	s.health.Category = category
}

// SetIngredients sets the ingredients of the SkinCare product.
func (s *SkinCareBuilder) SetIngredients(ingredients []string) {
	s.health.Ingredients = ingredients
}

// SetBrand sets the brand of the SkinCare product.
func (s *SkinCareBuilder) SetBrand(brand string) {
	s.health.Brand = brand
}

// SetUsageInstructions sets the usage instructions of the SkinCare product.
func (s *SkinCareBuilder) SetUsageInstructions(usageInstructions []string) {
	s.health.UsageInstructions = usageInstructions
}

// SetType sets the type of the SkinCare product.
func (s *SkinCareBuilder) SetType(healthType string) {
	s.health.Type = healthType
}

// SetSPF sets the SPF (Sun Protection Factor) of the SkinCare product.
func (s *SkinCareBuilder) SetSPF(spf string) {
	s.health.SPF = spf
}

// SetVolume sets the volume of the SkinCare product.
func (s *SkinCareBuilder) SetVolume(volume string) {
	s.health.Volume = volume
}

// SetForMenWomen sets whether the SkinCare product is for men or women.
func (s *SkinCareBuilder) SetForMenWomen(forMenWomen string) {
	s.health.ForMenWomen = forMenWomen
}

// GetItem returns the built SkinCare product.
func (s *SkinCareBuilder) GetItem() Health {
	return s.health
}

// Unused methods from abstract interface.

// SetColor sets the color of the SkinCare product.
func (s *SkinCareBuilder) SetColor(color string) {}

// SetFinish sets the finish of the SkinCare product.
func (s *SkinCareBuilder) SetFinish(finish string) {}

// SetSulfateFree sets whether the SkinCare product is sulfate-free.
func (s *SkinCareBuilder) SetSulfateFree(sulfateFree bool) {}

// SetSupplementForm sets the supplement form of the SkinCare product.
func (s *SkinCareBuilder) SetSupplementForm(supplementForm string) {}

// SetRecommendedDosage sets the recommended dosage of the SkinCare product.
func (s *SkinCareBuilder) SetRecommendedDosage(recommendedDosage string) {}
