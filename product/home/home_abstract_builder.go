package home

// HomeBuilder is an interface that defines the methods for building a home product.
type HomeBuilder interface {
	// SetAttributes sets the attributes of the home product which are common to all categories.
	SetAttributes(id, name string, price float64, description, vendor string)
	// SetCategory sets the category of the home product.
	SetCategory(category string)
	// SetBrand sets the brand of the home product.
	SetBrand(brand string)
	// SetMaterial sets the material of the home product.
	SetMaterial(material string)
	// SetDimensions sets the dimensions of the home product.
	SetDimensions(dimension string)
	// SetWeight sets the weight of the home product.
	SetWeight(weight string)
	// SetType sets the type of the home product.
	SetType(prodType string)
	// SetPowerSource sets the power source required for the home product.
	SetPowerSource(source string)
	// SetAssemblyRequired sets whether assembly is required for the home product.
	SetAssemblyRequired(required bool)
	// SetVolume sets the volume of the home product.
	SetVolume(volume string)
	// SetScent sets the scent of the home product.
	SetScent(scent string)
	// SetColor sets the color of the home product.
	SetColor(color string)
	// SetStyle sets the style of the home product.
	SetStyle(style string)
	// SetCapacity sets the capacity of the home product.
	SetCapacity(capacity string)
	// SetDishWasherSafe sets whether the home product is dishwasher safe.
	SetDishWasherSafe(sage bool)
	// Get the home product
	GetItem() Home
}

// getBuilder returns a HomeBuilder based on the provided builderType.
// It creates and returns a specific type of builder based on the input string.
// If the builderType is "bath", it returns a newBathBuilder.
// If the builderType is "tools", it returns a newToolBuilder.
// If the builderType is "kitchen", it returns a newKitchenBuilder.
// If the builderType is "furniture", it returns a newFurnitureBuilder.
// If the builderType is "decor", it returns a newDecorBuilder.
// If the builderType is not recognized, it returns nil.
func getBuilder(builderType string) HomeBuilder {
	if builderType == "bath" {
		return newBathBuilder()
	}
	if builderType == "tools" {
		return newToolBuilder()
	}
	if builderType == "kitchen" {
		return newKitchenBuilder()
	}
	if builderType == "furniture" {
		return newFurnitureBuilder()
	}
	if builderType == "decor" {
		return newDecorBuilder()
	}
	return nil

}
