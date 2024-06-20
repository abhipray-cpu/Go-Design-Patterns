package home

// ToolBuilder is a builder for creating Home objects with specific attributes.
type ToolBuilder struct {
	home Home
}

// newToolBuilder creates a new instance of ToolBuilder.
func newToolBuilder() *ToolBuilder {
	return &ToolBuilder{}
}

// SetAttributes sets the attributes of the Home object.
func (t *ToolBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	t.home.ID = id
	t.home.Name = name
	t.home.Price = price
	t.home.Description = description
	t.home.Vendor = vendor
}

// SetCategory sets the category of the Home object.
func (t *ToolBuilder) SetCategory(category string) {
	t.home.Category = category
}

// SetPrice sets the price of the Home object.
func (t *ToolBuilder) SetPrice(price float64) {
	t.home.Price = price
}

// SetBrand sets the brand of the Home object.
func (t *ToolBuilder) SetBrand(brand string) {
	t.home.Brand = brand
}

// SetType sets the type of the Home object.
func (t *ToolBuilder) SetType(toolType string) {
	t.home.Type = toolType
}

// SetColor sets the color of the Home object.
func (t *ToolBuilder) SetColor(color string) {
	t.home.Color = color
}

// SetMaterial sets the material of the Home object.
func (t *ToolBuilder) SetMaterial(material string) {
	t.home.Material = material
}

// SetWeight sets the weight of the Home object.
func (t *ToolBuilder) SetWeight(weight string) {
	t.home.Weight = weight
}

// SetStyle sets the style of the Home object.
func (t *ToolBuilder) SetStyle(style string) {
	t.home.Style = style
}

// SetDimensions sets the dimensions of the Home object.
func (t *ToolBuilder) SetDimensions(dimensions string) {
	t.home.Dimensions = dimensions
}

// SetPowerSource sets the power source of the Home object.
func (t *ToolBuilder) SetPowerSource(powerSource string) {
	t.home.PowerSource = powerSource
}

// GetItem returns the created Home object.
func (t *ToolBuilder) GetItem() Home {
	return t.home
}

// SetVolume is an unused method from the abstract interface.
func (t *ToolBuilder) SetVolume(volume string) {
}

// SetAssemblyRequired is an unused method from the abstract interface.
func (t *ToolBuilder) SetAssemblyRequired(required bool) {
}

// SetCapacity is an unused method from the abstract interface.
func (t *ToolBuilder) SetCapacity(capacity string) {
}

// SetDishWasherSafe is an unused method from the abstract interface.
func (t *ToolBuilder) SetDishWasherSafe(dishWasherSafe bool) {
}

// SetScent is an unused method from the abstract interface.
func (t *ToolBuilder) SetScent(scent string) {
}
