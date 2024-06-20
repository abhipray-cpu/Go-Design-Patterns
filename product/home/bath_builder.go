package home

// BathBuilder is a builder for creating Bath products.
type BathBuilder struct {
	home Home
}

// newBathBuilder creates a new instance of BathBuilder.
func newBathBuilder() *BathBuilder {
	return &BathBuilder{}
}

// SetAttributes sets the attributes of the Bath product.
func (b *BathBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	b.home.ID = id
	b.home.Name = name
	b.home.Price = price
	b.home.Description = description
	b.home.Vendor = vendor
}

// SetCategory sets the category of the Bath product.
func (b *BathBuilder) SetCategory(category string) {
	b.home.Category = category
}

// SetBrand sets the brand of the Bath product.
func (b *BathBuilder) SetBrand(brand string) {
	b.home.Brand = brand
}

// SetType sets the type of the Bath product.
func (b *BathBuilder) SetType(t string) {
	b.home.Type = t
}

// SetVolume sets the volume of the Bath product.
func (b *BathBuilder) SetVolume(volume string) {
	b.home.Volume = volume
}

// SetScent sets the scent of the Bath product.
func (b *BathBuilder) SetScent(scents string) {
	b.home.Scent = scents
}

// SetMaterial sets the material of the Bath product.
func (b *BathBuilder) SetMaterial(material string) {
	b.home.Material = material
}

// SetDimensions sets the dimensions of the Bath product.
func (b *BathBuilder) SetDimensions(dimensions string) {
	b.home.Dimensions = dimensions
}

// SetWeight sets the weight of the Bath product.
func (b *BathBuilder) SetWeight(weight string) {
	b.home.Weight = weight
}

// SetColor sets the color of the Bath product.
func (b *BathBuilder) SetColor(color string) {
	b.home.Color = color
}

// SetCapacity sets the capacity of the Bath product.
func (b *BathBuilder) SetCapacity(capacity string) {
	b.home.Capacity = capacity
}

// SetAssemblyRequired sets whether assembly is required for the Bath product.
func (b *BathBuilder) SetAssemblyRequired(required bool) {
	b.home.AssemblyRequired = required
}

// SetStyle sets the style of the Bath product.
func (b *BathBuilder) SetStyle(style string) {
	b.home.Style = style
}

// GetItem returns the built Bath product.
func (b *BathBuilder) GetItem() Home {
	return b.home
}

// unused methods from abstract interface

// SetDishWasherSafe sets whether the Bath product is dishwasher safe.
func (b *BathBuilder) SetDishWasherSafe(dishWasherSafe bool) {}

// SetPowerSource sets the power source of the Bath product.
func (b *BathBuilder) SetPowerSource(source string) {}
