package pet

// PetBuilder is an interface for building pet objects.
type PetBuilder interface {
	// sets the common attributes of all kinds of pet objects
	SetAttributes(id, name string, price float64, description, vendor string)
	// sets the category of the pet object
	SetCategory(category string)
	// sets the brand of the pet object
	SetBrand(brand string)
	// sets the ingredients of the pet food
	SetIngredients(ingredients []string)
	// sets the nutritional value of the pet food
	SetNutritionalValue(nutritionalValue string)
	// sets the weight of the pet objects
	SetWeight(weight string)
	// sets the suitable for which pets
	SetSuitableFor(suitable []string)
	// sets the type of the pet accessories
	SetType(prodType string)
	// sets the material of the pet objects
	SetMaterial(material string)
	// sets the size of the pet objects
	SetSize(size string)
	// sets the volume of the pet grooming items
	SetVolume(volume string)
	// sets the durability of the pet toys
	SetDurability(durability string)
	// GetSports returns the pets product.
	GetItem() Pet
}

// getBuilder returns a PetBuilder based on the given builderType.
// It takes a string parameter builderType and returns a PetBuilder interface.
// The builderType can be one of the following values: "food", "toy", "grooming", or "accessory".
// If the builderType is not recognized, it returns nil.
func getBuilder(builderType string) PetBuilder {
	switch builderType {
	case "food":
		return newFoodBuilder()
	case "toy":
		return newToyBuilder()
	case "grooming":
		return newGroomingBuilder()
	case "accessory":
		return newAccessoryBuilder()
	default:
		return nil
	}
}
