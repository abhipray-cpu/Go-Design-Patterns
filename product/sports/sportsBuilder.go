package sports

// SportsBuilder is an interface that defines the methods for building a sports product.
// SportsBuilder is an interface that defines the methods for building a sports product.
type SportsBuilder interface {
	// SetAttributes sets the attributes of the sports product.
	SetAttributes(id, name string, price float64, description string)

	// SetCategory sets the category of the sports product.
	SetCategory(category string)

	// SetAgeGroup sets the age group of the sports product.
	SetAgeGroup(minAge, maxAge int)

	// SetWeight sets the weight of the sports product.
	SetWeight(weight string)

	// set brand of the sports product
	SetBrand(brand string)
	
	// SetMaterial sets the material of the sports product.
	SetMaterial(material string)

	// SetDimensions sets the dimensions of the sports product.
	SetDimensions(dimensions string)

    // SetType sets the type of the sports product.
	SetType(prodType string)

	// SetDurability sets the durability of the sports product.
	SetDurability(durability string)
    
	// SetSuitableFor sets the suitable for of the sports product.
	SetSuitableFor(suitableFor []string)

	// SetSize sets the size of the sports product.
    SetSize(size string)

	// SetVendor sets the vendor of the sports product.
	SetVendor(vendor string)
	
	// SetColor of the sports product
	SetColor(color string)

	// GetSports returns the sports product.
	GetItem() Sports
}

// getBuilder returns a SportsBuilder based on the specified builderType.
// It accepts a string parameter builderType and returns a SportsBuilder interface.
// If the builderType is "Camping", it returns a new instance of CampingBuilder.
// If the builderType is "Equipment", it returns a new instance of EquipmentBuilder.
// If the builderType is "Exercise", it returns a new instance of ExerciseBuilder.
// If the builderType is "Recreation", it returns a new instance of RecreationBuilder.
// If the builderType is not recognized, it returns nil.
func getBuilder(builderType string) SportsBuilder {
	switch builderType {
	case "Camping":
		return newCampingBuilder()
	case "Equipment":
		return newEquipmentBuilder()
	case "Exercise":
		return newExerciseBuilder()
	case "Recreation":
		return newRecreationBuilder()
	default:
		return nil
	}
}