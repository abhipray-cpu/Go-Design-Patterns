package home

// KitchenBuilder is a builder for creating kitchen products.
type KitchenBuilder struct {
	home Home
}

// newKitchenBuilder creates a new instance of KitchenBuilder.
func newKitchenBuilder() *KitchenBuilder {
	return &KitchenBuilder{}
}

// SetAttributes sets the attributes of the kitchen product.
func (k *KitchenBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	k.home.ID = id
	k.home.Name = name
	k.home.Price = price
	k.home.Description = description
	k.home.Vendor = vendor
}

// SetCategory sets the category of the kitchen product.
func (k *KitchenBuilder) SetCategory(category string) {
	k.home.Category = category
}

// SetBrand sets the brand of the kitchen product.
func (k *KitchenBuilder) SetBrand(brand string) {
	k.home.Brand = brand
}

// SetType sets the type of the kitchen product.
func (k *KitchenBuilder) SetType(t string) {
	k.home.Type = t
}

// SetColor sets the color of the kitchen product.
func (k *KitchenBuilder) SetColor(color string) {
	k.home.Color = color
}

// SetMaterial sets the material of the kitchen product.
func (k *KitchenBuilder) SetMaterial(material string) {
	k.home.Material = material
}

// SetWeight sets the weight of the kitchen product.
func (k *KitchenBuilder) SetWeight(weight string) {
	k.home.Weight = weight
}

// SetStyle sets the style of the kitchen product.
func (k *KitchenBuilder) SetStyle(style string) {
	k.home.Style = style
}

// SetDimensions sets the dimensions of the kitchen product.
func (k *KitchenBuilder) SetDimensions(dimensions string) {
	k.home.Dimensions = dimensions
}

// SetCapacity sets the capacity of the kitchen product.
func (k *KitchenBuilder) SetCapacity(capacity string) {
	k.home.Capacity = capacity
}

// SetVolume sets the volume of the kitchen product.
func (k *KitchenBuilder) SetVolume(volume string) {
	k.home.Volume = volume
}

// SetDishWasherSafe sets whether the kitchen product is dishwasher safe.
func (k *KitchenBuilder) SetDishWasherSafe(dishWasherSafe bool) {
	k.home.DishWasherSafe = dishWasherSafe
}

// SetAssemblyRequired sets whether assembly is required for the kitchen product.
func (k *KitchenBuilder) SetAssemblyRequired(required bool) {
	k.home.AssemblyRequired = required
}

// SetPowerSource sets the power source of the kitchen product.
func (k *KitchenBuilder) SetPowerSource(powerSource string) {
	k.home.PowerSource = powerSource
}

// GetItem returns the built kitchen product.
func (k *KitchenBuilder) GetItem() Home {
	return k.home
}

// SetScent is an unused method from the abstract interface.
func (k *KitchenBuilder) SetScent(scents string) {
	// This method is not used in the KitchenBuilder implementation.
}
