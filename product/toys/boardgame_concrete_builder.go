package toys

// BoardToyBuilder is the concrete builder for board games that implements the ToyBuilder interface.
type BoardToyBuilder struct {
	toy Toy
}

// NewBoardToyBuilder creates a new instance of BoardToyBuilder.
func NewBoardToyBuilder() *BoardToyBuilder {
	return &BoardToyBuilder{toy: Toy{}}
}

// SetAttributes sets the attributes of the toy.
func (b *BoardToyBuilder) SetAttributes(id string, name string, price float64, description string) {
	b.toy.ID = id
	b.toy.Name = name
	b.toy.Price = price
	b.toy.Description = description
}

// SetCategory sets the category of the toy.
func (b *BoardToyBuilder) SetCategory(category string) {
	b.toy.Category = category
}

// SetAgeRange sets the age range of the toy.
func (b *BoardToyBuilder) SetAgeRange(minAge, maxAge int) {
	b.toy.MinAge = minAge
	b.toy.MaxAge = maxAge
}

// SetBrand sets the brand of the toy.
func (b *BoardToyBuilder) SetBrand(brand string) {
	b.toy.Brand = brand
}

// SetWeight sets the weight of the toy.
func (b *BoardToyBuilder) SetWeight(weight string) {
	b.toy.Weight = weight
}

// SetVendor sets the vendor of the toy.
func (b *BoardToyBuilder) SetVendor(vendor string) {
	b.toy.Vendor = vendor
}

// SetType sets the type of the toy.
func (b *BoardToyBuilder) SetType(prodType string) {
	b.toy.Type = prodType
}

// SetDimensions sets the dimensions of the toy.
func (b *BoardToyBuilder) SetDimensions(dimension string) {
	b.toy.Dimensions = dimension
}

// SetSuitableFor sets the suitable age range for the toy.
func (b *BoardToyBuilder) SetSuitableFor(suitableFor []string) {
	b.toy.Suitable = suitableFor
}

// GetToy returns the built toy.
func (b *BoardToyBuilder) GetToy() Toy {
	return b.toy
}

// SetMaterial is an unused function from the abstract interface.
func (b *BoardToyBuilder) SetMaterial(material string) {}

// SetColor is an unused function from the abstract interface.
func (b *BoardToyBuilder) SetColor(color string) {}

// SetSize is an unused function from the abstract interface.
func (b *BoardToyBuilder) SetSize(size string) {}
