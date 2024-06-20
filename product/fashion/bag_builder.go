package fashion

// BagBuilder is a builder for creating fashion bags.
type BagBuilder struct {
	fashion Fashion
}

// newBagBuilder creates a new instance of BagBuilder.
func newBagBuilder() *BagBuilder {
	return &BagBuilder{}
}

// SetAttributes sets the attributes of the fashion bag.
func (b *BagBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	b.fashion.Id = id
	b.fashion.Name = name
	b.fashion.Price = price
	b.fashion.Description = description
	b.fashion.Vendor = vendor
}

// SetBrand sets the brand of the fashion bag.
func (b *BagBuilder) SetBrand(brand string) {
	b.fashion.Brand = brand
}

// SetCategory sets the category of the fashion bag.
func (b *BagBuilder) SetCategory(category string) {
	b.fashion.Category = category
}

// SetMaterial sets the material of the fashion bag.
func (b *BagBuilder) SetMaterial(material string) {
	b.fashion.Material = material
}

// SetSize sets the size of the fashion bag.
func (b *BagBuilder) SetSize(size string) {
	b.fashion.Size = size
}

// SetColor sets the color of the fashion bag.
func (b *BagBuilder) SetColor(color string) {
	b.fashion.Color = color
}

// SetGender sets the gender of the fashion bag.
func (b *BagBuilder) SetGender(gender string) {
	b.fashion.Gender = gender
}

// SetStyle sets the style of the fashion bag.
func (b *BagBuilder) SetStyle(style string) {
	b.fashion.Style = style
}

// SetType sets the type of the fashion bag.
func (b *BagBuilder) SetType(fashionType string) {
	b.fashion.Type = fashionType
}

// GetItem returns the created fashion bag.
func (b *BagBuilder) GetItem() Fashion {
	return b.fashion
}

// SetGemStone is an unused method in the BagBuilder interface.
func (b *BagBuilder) SetGemStone(gemStone string) {}

// SetWaterResistance is an unused method in the BagBuilder interface.
func (b *BagBuilder) SetWaterResistance(waterResistance string) {}
