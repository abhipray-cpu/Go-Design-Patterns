package sports

// CampingBuilder is a builder for creating camping products.
type CampingBuilder struct {
	sports Sports
}

// newCampingBuilder creates a new instance of CampingBuilder.
func newCampingBuilder() *CampingBuilder {
	return &CampingBuilder{}
}

// SetAttributes sets the attributes of the camping product.
func (c *CampingBuilder) SetAttributes(id string, name string, price float64, description string) {
	c.sports.ID = id
	c.sports.Name = name
	c.sports.Price = price
	c.sports.Description = description
}

// SetCategory sets the category of the camping product.
func (c *CampingBuilder) SetCategory(category string) {
	c.sports.Category = category
}

// SetAgeGroup sets the age group for which the camping product is suitable.
func (c *CampingBuilder) SetAgeGroup(minAge, maxAge int) {
	c.sports.MinAge = minAge
	c.sports.MaxAge = maxAge
}

// SetWeight sets the weight of the camping product.
func (c *CampingBuilder) SetWeight(weight string) {
	c.sports.Weight = weight
}

// SetMaterial sets the material of the camping product.
func (c *CampingBuilder) SetMaterial(material string) {
	c.sports.Material = material
}

// SetDimensions sets the dimensions of the camping product.
func (c *CampingBuilder) SetDimensions(dimension string) {
	c.sports.Dimensions = dimension
}

// SetDurability sets the durability of the camping product.
func (c *CampingBuilder) SetDurability(durability string) {
	c.sports.Durability = durability
}

// SetVendor sets the vendor of the camping product.
func (c *CampingBuilder) SetVendor(vendor string) {
	c.sports.Vendor = vendor
}

// SetType sets the type of the camping product.
func (c *CampingBuilder) SetType(prodType string) {
	c.sports.Type = prodType
}

// SetSuitableFor sets the suitable age groups for the camping product.
func (c *CampingBuilder) SetSuitableFor(suitableFor []string) {
	c.sports.Suitable = suitableFor
}

// SetSize sets the size of the camping product.
func (c *CampingBuilder) SetSize(size string) {
	c.sports.Size = size
}

// SetBrand sets the brand of the camping product.
func (c *CampingBuilder) SetBrand(brand string) {
	c.sports.Brand = brand
}

// SetColor sets the color of the camping product.
func (c *CampingBuilder) SetColor(color string) {
	c.sports.Color = color
}

// GetItem returns the built camping product.
func (c *CampingBuilder) GetItem() Sports {
	return c.sports
}

// unused methods from abstract interface
