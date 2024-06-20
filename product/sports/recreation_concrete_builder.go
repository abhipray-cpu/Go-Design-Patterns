package sports

// RecreationBuilder is a builder for creating recreation sports products.
type RecreationBuilder struct {
	sports Sports
}

// newRecreationBuilder creates a new instance of RecreationBuilder.
func newRecreationBuilder() *RecreationBuilder {
	return &RecreationBuilder{}
}

// SetAttributes sets the attributes of the recreation sports product.
func (r *RecreationBuilder) SetAttributes(id string, name string, price float64, description string) {
	r.sports.ID = id
	r.sports.Name = name
	r.sports.Price = price
	r.sports.Description = description
}

// SetCategory sets the category of the recreation sports product.
func (r *RecreationBuilder) SetCategory(category string) {
	r.sports.Category = category
}

// SetAgeGroup sets the age group of the recreation sports product.
func (r *RecreationBuilder) SetAgeGroup(minAge, maxAge int) {
	r.sports.MinAge = minAge
	r.sports.MaxAge = maxAge
}

// SetWeight sets the weight of the recreation sports product.
func (r *RecreationBuilder) SetWeight(weight string) {
	r.sports.Weight = weight
}

// SetMaterial sets the material of the recreation sports product.
func (r *RecreationBuilder) SetMaterial(material string) {
	r.sports.Material = material
}

// SetDimensions sets the dimensions of the recreation sports product.
func (r *RecreationBuilder) SetDimensions(dimension string) {
	r.sports.Dimensions = dimension
}

// SetDurability sets the durability of the recreation sports product.
func (r *RecreationBuilder) SetDurability(durability string) {
	r.sports.Durability = durability
}

// SetVendor sets the vendor of the recreation sports product.
func (r *RecreationBuilder) SetVendor(vendor string) {
	r.sports.Vendor = vendor
}

// SetType sets the type of the recreation sports product.
func (r *RecreationBuilder) SetType(prodType string) {
	r.sports.Type = prodType
}

// SetSuitableFor sets the suitable for field of the recreation sports product.
func (r *RecreationBuilder) SetSuitableFor(suitableFor []string) {
	r.sports.Suitable = suitableFor
}

// SetSize sets the size of the recreation sports product.
func (r *RecreationBuilder) SetSize(size string) {
	r.sports.Size = size
}

// SetColor sets the color of the recreation sports product.
func (r *RecreationBuilder) SetColor(color string) {
	r.sports.Color = color
}

// SetBrand sets the brand of the recreation sports product.
func (r *RecreationBuilder) SetBrand(brand string) {
	r.sports.Brand = brand
}

// GetItem returns the recreation sports product.
func (r *RecreationBuilder) GetItem() Sports {
	return r.sports
}

// Unused functions from abstract interface.
