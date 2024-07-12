package health

// SupplementBuilder is a builder for creating health supplements.
type SupplementBuilder struct {
	health Health
}

// newSupplementBuilder creates a new instance of SupplementBuilder.
func newSupplementBuilder() *SupplementBuilder {
	return &SupplementBuilder{}
}

// SetAttributes sets the attributes of the health supplement.
func (s *SupplementBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	s.health.Id = id
	s.health.Name = name
	s.health.Price = price
	s.health.Description = description
	s.health.Vendor = vendor
}

// SetCategory sets the category of the health supplement.
func (s *SupplementBuilder) SetCategory(category string) {
	s.health.Category = category
}

// SetIngredients sets the ingredients of the health supplement.
func (s *SupplementBuilder) SetIngredients(ingredients []string) {
	s.health.Ingredients = ingredients
}

// SetBrand sets the brand of the health supplement.
func (s *SupplementBuilder) SetBrand(brand string) {
	s.health.Brand = brand
}

// SetType sets the type of the health supplement.
func (s *SupplementBuilder) SetType(healthType string) {
	s.health.Type = healthType
}

// SetVolume sets the volume of the health supplement.
func (s *SupplementBuilder) SetVolume(volume string) {
	s.health.Volume = volume
}

// SetSupplementForm sets the form of the health supplement.
func (s *SupplementBuilder) SetSupplementForm(supplementForm string) {
	s.health.SupplementForm = supplementForm
}

// SetRecommendedDosage sets the recommended dosage of the health supplement.
func (s *SupplementBuilder) SetRecommendedDosage(recommendedDosage string) {
	s.health.RecommendedDosage = recommendedDosage
}

// SetForMenWomen sets the target audience (men/women) of the health supplement.
func (s *SupplementBuilder) SetForMenWomen(forMenWomen string) {
	s.health.ForMenWomen = forMenWomen
}

// SetUsageInstructions sets the usage instructions of the health supplement.
func (s *SupplementBuilder) SetUsageInstructions(usageInstructions []string) {
	s.health.UsageInstructions = usageInstructions
}

// GetItem returns the built Health object.
func (s *SupplementBuilder) GetItem() Health {
	return s.health
}

// unused exposef methods from abstract interface

// SetSPF sets the SPF (Sun Protection Factor) of the health supplement.
func (s *SupplementBuilder) SetSPF(spf string) {
}

// SetColor sets the color of the health supplement.
func (s *SupplementBuilder) SetColor(color string) {
}

// SetFinish sets the finish of the health supplement.
func (s *SupplementBuilder) SetFinish(finish string) {
}

// SetCoverage sets the coverage of the health supplement.
func (s *SupplementBuilder) SetCoverage(coverage string) {
}

// SetSulfateFree sets whether the health supplement is sulfate-free.
func (s *SupplementBuilder) SetSulfateFree(sulfateFree bool) {
}
