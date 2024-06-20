package fashion

// ShoesBuilder is a builder for creating shoes fashion items.
type ShoesBuilder struct {
	fashion Fashion
}

// newShoesBuilder creates a new instance of ShoesBuilder.
func newShoesBuilder() *ShoesBuilder {
	return &ShoesBuilder{}
}

// SetAttributes sets the attributes of the shoes fashion item.
func (s *ShoesBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	s.fashion.Id = id
	s.fashion.Name = name
	s.fashion.Price = price
	s.fashion.Description = description
	s.fashion.Vendor = vendor
}

// SetBrand sets the brand of the shoes fashion item.
func (s *ShoesBuilder) SetBrand(brand string) {
	s.fashion.Brand = brand
}

// SetCategory sets the category of the shoes fashion item.
func (s *ShoesBuilder) SetCategory(category string) {
	s.fashion.Category = category
}

// SetMaterial sets the material of the shoes fashion item.
func (s *ShoesBuilder) SetMaterial(material string) {
	s.fashion.Material = material
}

// SetSize sets the size of the shoes fashion item.
func (s *ShoesBuilder) SetSize(size string) {
	s.fashion.Size = size
}

// SetColor sets the color of the shoes fashion item.
func (s *ShoesBuilder) SetColor(color string) {
	s.fashion.Color = color
}

// SetGender sets the gender of the shoes fashion item.
func (s *ShoesBuilder) SetGender(gender string) {
	s.fashion.Gender = gender
}

// SetStyle sets the style of the shoes fashion item.
func (s *ShoesBuilder) SetStyle(style string) {
	s.fashion.Style = style
}

// SetType sets the type of the shoes fashion item.
func (s *ShoesBuilder) SetType(fashionType string) {
	s.fashion.Type = fashionType
}

// GetItem returns the created shoes fashion item.
func (s *ShoesBuilder) GetItem() Fashion {
	return s.fashion
}

// SetWaterResistance sets the water resistance of the shoes fashion item.
// This method is not used in the current implementation.
func (s *ShoesBuilder) SetWaterResistance(waterResistance string) {
	// Not implemented
}

// SetGemStone sets the gemstone of the shoes fashion item.
// This method is not used in the current implementation.
func (s *ShoesBuilder) SetGemStone(gemStone string) {
	// Not implemented
}
