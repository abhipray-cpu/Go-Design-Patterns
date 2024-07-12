package toys

// ActionToyBuilder is the concrete builder for action toys that implements the ToyBuilder interface.
type ActionToyBuilder struct {
	toy Toy
}

// NewActionToyBuilder creates a new instance of ActionToyBuilder.
func NewActionToyBuilder() *ActionToyBuilder {
	return &ActionToyBuilder{toy: Toy{}}
}

// SetAttributes sets the attributes of the action toy.
func (a *ActionToyBuilder) SetAttributes(id string, name string, price float64, description string) {
	a.toy.ID = id
	a.toy.Name = name
	a.toy.Price = price
	a.toy.Description = description
}

// SetCategory sets the category of the action toy.
func (a *ActionToyBuilder) SetCategory(category string) {
	a.toy.Category = category
}

// SetAgeRange sets the age range of the action toy.
func (a *ActionToyBuilder) SetAgeRange(minAge, maxAge int) {
	a.toy.MinAge = minAge
	a.toy.MaxAge = maxAge
}

// SetBrand sets the brand of the action toy.
func (a *ActionToyBuilder) SetBrand(brand string) {
	a.toy.Brand = brand
}

// SetWeight sets the weight of the action toy.
func (a *ActionToyBuilder) SetWeight(weight string) {
	a.toy.Weight = weight
}

// SetMaterial sets the material of the action toy.
func (a *ActionToyBuilder) SetMaterial(material string) {
	a.toy.Material = material
}

// SetColor sets the color of the action toy.
func (a *ActionToyBuilder) SetColor(color string) {
	a.toy.Color = color
}

// SetVendor sets the vendor of the action toy.
func (a *ActionToyBuilder) SetVendor(vendor string) {
	a.toy.Vendor = vendor
}

// SetSize sets the size of the action toy.
func (a *ActionToyBuilder) SetSize(size string) {
	a.toy.Size = size
}

// SetType sets the type of the action toy.
func (a *ActionToyBuilder) SetType(prodType string) {
	a.toy.Type = prodType
}

// SetSuitableFor sets the suitable age range for the action toy.
func (a *ActionToyBuilder) SetSuitableFor(suitableFor []string) {
	a.toy.Suitable = suitableFor
}

// SetDimensions sets the dimensions of the action toy.
func (a *ActionToyBuilder) SetDimensions(dimension string) {
	a.toy.Dimensions = dimension
}

// GetToy returns the built action toy.
func (a *ActionToyBuilder) GetToy() Toy {
	return a.toy
}

// Unused functions from the abstract interface.
