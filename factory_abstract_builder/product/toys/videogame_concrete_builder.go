package toys

// VideoToyBuilder is the concrete builder for video games that implements the ToyBuilder interface.
type VideoToyBuilder struct {
	toy Toy
}

// NewVideoToyBuilder creates a new instance of VideoToyBuilder.
func NewVideoToyBuilder() *VideoToyBuilder {
	return &VideoToyBuilder{toy: Toy{}}
}

// SetAttributes sets the attributes of the toy.
func (v *VideoToyBuilder) SetAttributes(id string, name string, price float64, description string) {
	v.toy.ID = id
	v.toy.Name = name
	v.toy.Price = price
	v.toy.Description = description
}

// SetCategory sets the category of the toy.
func (v *VideoToyBuilder) SetCategory(category string) {
	v.toy.Category = category
}

// SetAgeRange sets the age range of the toy.
func (v *VideoToyBuilder) SetAgeRange(minAge, maxAge int) {
	v.toy.MinAge = minAge
	v.toy.MaxAge = maxAge
}

// SetBrand sets the brand of the toy.
func (v *VideoToyBuilder) SetBrand(brand string) {
	v.toy.Brand = brand
}

// SetWeight sets the weight of the toy.
func (v *VideoToyBuilder) SetWeight(weight string) {
	v.toy.Weight = weight
}

// SetType sets the type of the toy.
func (v *VideoToyBuilder) SetType(prodType string) {
	v.toy.Type = prodType
}

// SetVendor sets the vendor of the toy.
func (v *VideoToyBuilder) SetVendor(vendor string) {
	v.toy.Vendor = vendor
}

// SetSuitableFor sets the suitable age range for the toy.
func (v *VideoToyBuilder) SetSuitableFor(suitableFor []string) {
	v.toy.Suitable = suitableFor
}

// GetToy returns the built toy.
func (v *VideoToyBuilder) GetToy() Toy {
	return v.toy
}

// SetMaterial sets the material of the toy. (Unused)
func (v *VideoToyBuilder) SetMaterial(material string) {
	v.toy.Material = material
}

// SetColor sets the color of the toy. (Unused)
func (v *VideoToyBuilder) SetColor(color string) {
	v.toy.Color = color
}

// SetSize sets the size of the toy. (Unused)
func (v *VideoToyBuilder) SetSize(size string) {
	// This function is not implemented.
}

// SetDimensions sets the dimensions of the toy. (Unused)
func (v *VideoToyBuilder) SetDimensions(dimension string) {
	// This function is not implemented.
}
