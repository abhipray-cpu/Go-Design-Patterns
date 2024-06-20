package home

// DecorBuilder is a builder for creating Home objects with specific attributes.
type DecorBuilder struct {
	home Home
}

// newDecorBuilder creates a new instance of DecorBuilder.
func newDecorBuilder() *DecorBuilder {
	return &DecorBuilder{}
}

// SetAttributes sets the attributes of the Home object.
func (d *DecorBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	d.home.ID = id
	d.home.Name = name
	d.home.Price = price
	d.home.Description = description
	d.home.Vendor = vendor
}

// SetCategory sets the category of the Home object.
func (d *DecorBuilder) SetCategory(category string) {
	d.home.Category = category
}

// SetBrand sets the brand of the Home object.
func (d *DecorBuilder) SetBrand(brand string) {
	d.home.Brand = brand
}

// SetType sets the type of the Home object.
func (d *DecorBuilder) SetType(t string) {
	d.home.Type = t
}

// SetColor sets the color of the Home object.
func (d *DecorBuilder) SetColor(color string) {
	d.home.Color = color
}

// SetMaterial sets the material of the Home object.
func (d *DecorBuilder) SetMaterial(material string) {
	d.home.Material = material
}

// SetWeight sets the weight of the Home object.
func (d *DecorBuilder) SetWeight(weight string) {
	d.home.Weight = weight
}

// SetStyle sets the style of the Home object.
func (d *DecorBuilder) SetStyle(style string) {
	d.home.Style = style
}

// SetDimensions sets the dimensions of the Home object.
func (d *DecorBuilder) SetDimensions(dimensions string) {
	d.home.Dimensions = dimensions
}

// SetAssemblyRequired sets whether assembly is required for the Home object.
func (d *DecorBuilder) SetAssemblyRequired(required bool) {
	d.home.AssemblyRequired = required
}

// GetItem returns the created Home object.
func (d *DecorBuilder) GetItem() Home {
	return d.home
}

// unused methods from abstract interface

// SetScent sets the scent of the Home object.
func (d *DecorBuilder) SetScent(scents string) {}

// SetVolume sets the volume of the Home object.
func (d *DecorBuilder) SetVolume(volume string) {}

// SetCapacity sets the capacity of the Home object.
func (d *DecorBuilder) SetCapacity(capacity string) {}

// SetPowerSource sets the power source of the Home object.
func (d *DecorBuilder) SetPowerSource(powerSource string) {}

// SetDishWasherSafe sets whether the Home object is dishwasher safe.
func (d *DecorBuilder) SetDishWasherSafe(dishWasherSafe bool) {}
