package fashion

// WatchBuilder is a builder for creating Watch objects.
type WatchBuilder struct {
	fashion Fashion
}

// newWatchBuilder creates a new instance of WatchBuilder.
func newWatchBuilder() *WatchBuilder {
	return &WatchBuilder{}
}

// SetAttributes sets the attributes of the Watch.
func (w *WatchBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	w.fashion.Id = id
	w.fashion.Name = name
	w.fashion.Price = price
	w.fashion.Description = description
	w.fashion.Vendor = vendor
}

// SetBrand sets the brand of the Watch.
func (w *WatchBuilder) SetBrand(brand string) {
	w.fashion.Brand = brand
}

// SetCategory sets the category of the Watch.
func (w *WatchBuilder) SetCategory(category string) {
	w.fashion.Category = category
}

// SetMaterial sets the material of the Watch.
func (w *WatchBuilder) SetMaterial(material string) {
	w.fashion.Material = material
}

// SetSize sets the size of the Watch.
func (w *WatchBuilder) SetSize(size string) {
	w.fashion.Size = size
}

// SetColor sets the color of the Watch.
func (w *WatchBuilder) SetColor(color string) {
	w.fashion.Color = color
}

// SetGender sets the gender of the Watch.
func (w *WatchBuilder) SetGender(gender string) {
	w.fashion.Gender = gender
}

// SetStyle sets the style of the Watch.
func (w *WatchBuilder) SetStyle(style string) {
	w.fashion.Style = style
}

// SetType sets the type of the Watch.
func (w *WatchBuilder) SetType(fashionType string) {
	w.fashion.Type = fashionType
}

// SetWaterResistance sets the water resistance of the Watch.
func (w *WatchBuilder) SetWaterResistance(waterResistance string) {
	w.fashion.WaterResistance = waterResistance
}

// GetItem returns the built Watch object.
func (w *WatchBuilder) GetItem() Fashion {
	return w.fashion
}

// SetGemStone is an unused method from the abstract interface.
func (w *WatchBuilder) SetGemStone(gemStone string) {
}
