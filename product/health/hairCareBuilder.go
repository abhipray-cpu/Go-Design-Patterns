package health

// HairCareBuilder is a builder for creating HairCare products.
type HairCareBuilder struct {
	health Health
}

// newHairCareBuilder creates a new instance of HairCareBuilder.
func newHairCareBuilder() *HairCareBuilder {
	return &HairCareBuilder{}
}

// SetAttributes sets the attributes of the HairCare product.
func (h *HairCareBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	h.health.Id = id
	h.health.Name = name
	h.health.Price = price
	h.health.Description = description
	h.health.Vendor = vendor
}

// SetCategory sets the category of the HairCare product.
func (h *HairCareBuilder) SetCategory(category string) {
	h.health.Category = category
}

// SetIngredients sets the ingredients of the HairCare product.
func (h *HairCareBuilder) SetIngredients(ingredients []string) {
	h.health.Ingredients = ingredients
}

// SetBrand sets the brand of the HairCare product.
func (h *HairCareBuilder) SetBrand(brand string) {
	h.health.Brand = brand
}

// SetUsageInstructions sets the usage instructions of the HairCare product.
func (h *HairCareBuilder) SetUsageInstructions(usageInstructions []string) {
	h.health.UsageInstructions = usageInstructions
}

// SetType sets the type of the HairCare product.
func (h *HairCareBuilder) SetType(healthType string) {
	h.health.Type = healthType
}

// SetSPF sets the SPF (Sun Protection Factor) of the HairCare product.
func (h *HairCareBuilder) SetSPF(spf string) {
	h.health.SPF = spf
}

// SetVolume sets the volume of the HairCare product.
func (h *HairCareBuilder) SetVolume(volume string) {
	h.health.Volume = volume
}

// SetSulfateFree sets whether the HairCare product is sulfate-free or not.
func (h *HairCareBuilder) SetSulfateFree(sulfateFree bool) {
	h.health.SulfateFree = sulfateFree
}

// SetForMenWomen sets whether the HairCare product is for men or women.
func (h *HairCareBuilder) SetForMenWomen(forMenWomen string) {
	h.health.ForMenWomen = forMenWomen
}

// GetItem returns the built HairCare product.
func (h *HairCareBuilder) GetItem() Health {
	return h.health
}

// The following functions are unused in the abstract interface.

// SetColor sets the color of the HairCare product.
func (h *HairCareBuilder) SetColor(color string) {}

// SetFinish sets the finish of the HairCare product.
func (h *HairCareBuilder) SetFinish(finish string) {}

// SetSupplementForm sets the supplement form of the HairCare product.
func (h *HairCareBuilder) SetSupplementForm(supplementForm string) {}

// SetRecommendedDosage sets the recommended dosage of the HairCare product.
func (h *HairCareBuilder) SetRecommendedDosage(recommendedDosage string) {}
