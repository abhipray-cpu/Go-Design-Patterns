package fashion

// ClothingBuilder is a builder for creating clothing objects.
type ClothingBuilder struct {
	fashion Fashion
}

// newClothingBuilder creates a new instance of ClothingBuilder.
func newClothingBuilder() *ClothingBuilder {
	return &ClothingBuilder{}
}

// SetAttributes sets the attributes of the clothing.
func (c *ClothingBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	c.fashion.Id = id
	c.fashion.Name = name
	c.fashion.Price = price
	c.fashion.Description = description
	c.fashion.Vendor = vendor
}

// SetCategory sets the category of the clothing.
func (c *ClothingBuilder) SetCategory(category string) {
	c.fashion.Category = category
}

// SetBrand sets the brand of the clothing.
func (c *ClothingBuilder) SetBrand(brand string) {
	c.fashion.Brand = brand
}

// SetMaterial sets the material of the clothing.
func (c *ClothingBuilder) SetMaterial(material string) {
	c.fashion.Material = material
}

// SetSize sets the size of the clothing.
func (c *ClothingBuilder) SetSize(size string) {
	c.fashion.Size = size
}

// SetColor sets the color of the clothing.
func (c *ClothingBuilder) SetColor(color string) {
	c.fashion.Color = color
}

// SetGender sets the gender of the clothing.
func (c *ClothingBuilder) SetGender(gender string) {
	c.fashion.Gender = gender
}

// SetStyle sets the style of the clothing.
func (c *ClothingBuilder) SetStyle(style string) {
	c.fashion.Style = style
}

// SetType sets the type of the clothing.
func (c *ClothingBuilder) SetType(fashionType string) {
	c.fashion.Type = fashionType
}

// GetItem returns the created Fashion object.
func (c *ClothingBuilder) GetItem() Fashion {
	return c.fashion
}

// SetGemStone is an unused method from the abstract interface.
func (c *ClothingBuilder) SetGemStone(gemStone string) {
}

// SetWaterResistance is an unused method from the abstract interface.
func (c *ClothingBuilder) SetWaterResistance(waterResistance string) {
}
