package sports

// EquipmentBuilder is a builder for creating sports equipment.
type EquipmentBuilder struct {
	sports Sports
}

// newEquipmentBuilder creates a new instance of EquipmentBuilder.
func newEquipmentBuilder() *EquipmentBuilder {
	return &EquipmentBuilder{}
}

// SetAttributes sets the attributes of the sports equipment.
func (e *EquipmentBuilder) SetAttributes(id string, name string, price float64, description string) {
	e.sports.ID = id
	e.sports.Name = name
	e.sports.Price = price
	e.sports.Description = description
}

// SetCategory sets the category of the sports equipment.
func (e *EquipmentBuilder) SetCategory(category string) {
	e.sports.Category = category
}

// SetAgeGroup sets the age group for which the sports equipment is suitable.
func (e *EquipmentBuilder) SetAgeGroup(minAge, maxAge int) {
	e.sports.MinAge = minAge
	e.sports.MaxAge = maxAge
}

// SetWeight sets the weight of the sports equipment.
func (e *EquipmentBuilder) SetWeight(weight string) {
	e.sports.Weight = weight
}

// SetMaterial sets the material of the sports equipment.
func (e *EquipmentBuilder) SetMaterial(material string) {
	e.sports.Material = material
}

// SetDimensions sets the dimensions of the sports equipment.
func (e *EquipmentBuilder) SetDimensions(dimension string) {
	e.sports.Dimensions = dimension
}

// SetDurability sets the durability of the sports equipment.
func (e *EquipmentBuilder) SetDurability(durability string) {
	e.sports.Durability = durability
}

// SetVendor sets the vendor of the sports equipment.
func (e *EquipmentBuilder) SetVendor(vendor string) {
	e.sports.Vendor = vendor
}

// SetType sets the type of the sports equipment.
func (e *EquipmentBuilder) SetType(prodType string) {
	e.sports.Type = prodType
}

// SetSuitableFor sets the target audience for which the sports equipment is suitable.
func (e *EquipmentBuilder) SetSuitableFor(suitableFor []string) {
	e.sports.Suitable = suitableFor
}

// SetSize sets the size of the sports equipment.
func (e *EquipmentBuilder) SetSize(size string) {
	e.sports.Size = size
}

// SetColor sets the color of the sports equipment.
func (e *EquipmentBuilder) SetColor(color string) {
	e.sports.Color = color
}

// SetBrand sets the brand of the sports equipment.
func (e *EquipmentBuilder) SetBrand(brand string) {
	e.sports.Brand = brand
}

// GetItem returns the created sports equipment.
func (e *EquipmentBuilder) GetItem() Sports {
	return e.sports
}

// unuses methods from abstract interface
