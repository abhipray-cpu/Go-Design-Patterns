package home

// FurnitureBuilder is a builder for creating Furniture objects.
type FurnitureBuilder struct {
	home Home
}

// newFurnitureBuilder creates a new instance of FurnitureBuilder.
func newFurnitureBuilder() *FurnitureBuilder {
	return &FurnitureBuilder{}
}

// SetAttributes sets the attributes of the Furniture.
func (f *FurnitureBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	f.home.ID = id
	f.home.Name = name
	f.home.Price = price
	f.home.Description = description
	f.home.Vendor = vendor
}

// SetCategory sets the category of the Furniture.
func (f *FurnitureBuilder) SetCategory(category string) {
	f.home.Category = category
}

// SetBrand sets the brand of the Furniture.
func (f *FurnitureBuilder) SetBrand(brand string) {
	f.home.Brand = brand
}

// SetType sets the type of the Furniture.
func (f *FurnitureBuilder) SetType(t string) {
	f.home.Type = t
}

// SetColor sets the color of the Furniture.
func (f *FurnitureBuilder) SetColor(color string) {
	f.home.Color = color
}

// SetMaterial sets the material of the Furniture.
func (f *FurnitureBuilder) SetMaterial(material string) {
	f.home.Material = material
}

// SetWeight sets the weight of the Furniture.
func (f *FurnitureBuilder) SetWeight(weight string) {
	f.home.Weight = weight
}

// SetStyle sets the style of the Furniture.
func (f *FurnitureBuilder) SetStyle(style string) {
	f.home.Style = style
}

// SetDimensions sets the dimensions of the Furniture.
func (f *FurnitureBuilder) SetDimensions(dimensions string) {
	f.home.Dimensions = dimensions
}

// SetAssemblyRequired sets whether the Furniture requires assembly.
func (f *FurnitureBuilder) SetAssemblyRequired(required bool) {
	f.home.AssemblyRequired = required
}

// SetCapacity sets the capacity of the Furniture.
func (f *FurnitureBuilder) SetCapacity(capacity string) {
	f.home.Capacity = capacity
}

// GetItem returns the created Furniture object.
func (f *FurnitureBuilder) GetItem() Home {
	return f.home
}

// unused methods from abstract interface

// SetVolume sets the volume of the Furniture.
func (f *FurnitureBuilder) SetVolume(volume string) {}

// SetPowerSource sets the power source of the Furniture.
func (f *FurnitureBuilder) SetPowerSource(powerSource string) {}

// SetDishWasherSafe sets whether the Furniture is dishwasher safe.
func (f *FurnitureBuilder) SetDishWasherSafe(dishWasherSafe bool) {}

// SetScent sets the scent of the Furniture.
func (f *FurnitureBuilder) SetScent(scents string) {}
