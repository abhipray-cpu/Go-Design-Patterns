package fashion

// JewelryBuilder is a builder for creating jewelry fashion items.
type JewelryBuilder struct {
	fashion Fashion
}

// newJewelryBuilder creates a new instance of JewelryBuilder.
func newJewelryBuilder() *JewelryBuilder {
	return &JewelryBuilder{}
}

// SetAttributes sets the attributes of the jewelry fashion item.
func (j *JewelryBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	j.fashion.Id = id
	j.fashion.Name = name
	j.fashion.Price = price
	j.fashion.Description = description
	j.fashion.Vendor = vendor
}

// SetCategory sets the category of the jewelry fashion item.
func (j *JewelryBuilder) SetCategory(category string) {
	j.fashion.Category = category
}

// SetBrand sets the brand of the jewelry fashion item.
func (j *JewelryBuilder) SetBrand(brand string) {
	j.fashion.Brand = brand
}

// SetMaterial sets the material of the jewelry fashion item.
func (j *JewelryBuilder) SetMaterial(material string) {
	j.fashion.Material = material
}

// SetSize sets the size of the jewelry fashion item.
func (j *JewelryBuilder) SetSize(size string) {
	j.fashion.Size = size
}

// SetColor sets the color of the jewelry fashion item.
func (j *JewelryBuilder) SetColor(color string) {
	j.fashion.Color = color
}

// SetGender sets the gender of the jewelry fashion item.
func (j *JewelryBuilder) SetGender(gender string) {
	j.fashion.Gender = gender
}

// SetStyle sets the style of the jewelry fashion item.
func (j *JewelryBuilder) SetStyle(style string) {
	j.fashion.Style = style
}

// SetType sets the type of the jewelry fashion item.
func (j *JewelryBuilder) SetType(fashionType string) {
	j.fashion.Type = fashionType
}

// SetGemStone sets the gemstone of the jewelry fashion item.
func (j *JewelryBuilder) SetGemStone(gemStone string) {
	j.fashion.GemStone = gemStone
}

// GetItem returns the created jewelry fashion item.
func (j *JewelryBuilder) GetItem() Fashion {
	return j.fashion
}

// SetWaterResistance is an unused method from the abstract interface.
func (j *JewelryBuilder) SetWaterResistance(waterResistance string) {}
