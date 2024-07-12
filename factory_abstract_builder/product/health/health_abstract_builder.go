package health

// HealthBuilder is a struct that represents a builder for creating Health objects.
// HealthBuilder is an interface that defines the methods for building a health product.
type HealthBuilder interface {
	// SetAttributes sets the attributes of the health product.
	// It takes the ID, name, price, description, and vendor as parameters.
	SetAttributes(id, name string, price float64, description, vendor string)

	// SetCategory sets the category of the health product.
	// It takes the category as a parameter.
	SetCategory(category string)

	// SetIngredients sets the ingredients of the health product.
	// It takes the ingredients as a slice of strings.
	SetIngredients(ingredients []string)

	// SetBrand sets the brand of the health product.
	// It takes the brand as a parameter.
	SetBrand(brand string)

	// SetUsageInstructions sets the usage instructions of the health product.
	// It takes the usage instructions as a slice of strings.
	SetUsageInstructions(usageInstructions []string)

	// SetType sets the type of the health product.
	// It takes the health type as a parameter.
	SetType(healthType string)

	// SetSPF sets the SPF (Sun Protection Factor) of the health product.
	// It takes the SPF as a parameter.
	SetSPF(spf string)

	// SetVolume sets the volume of the health product.
	// It takes the volume as a parameter.
	SetVolume(volume string)

	// SetColor sets the color of the health product.
	// It takes the color as a parameter.
	SetColor(color string)

	// SetFinish sets the finish of the health product.
	// It takes the finish as a parameter.
	SetFinish(finish string)

	// SetSulfateFree sets whether the health product is sulfate-free or not.
	// It takes a boolean value indicating sulfate-free status.
	SetSulfateFree(sulfateFree bool)

	// SetForMenWomen sets whether the health product is for men, women, or both.
	// It takes the target audience as a parameter.
	SetForMenWomen(forMenWomen string)

	// SetSupplementForm sets the supplement form of the health product.
	// It takes the supplement form as a parameter.
	SetSupplementForm(supplementForm string)

	// SetRecommendedDosage sets the recommended dosage of the health product.
	// It takes the recommended dosage as a parameter.
	SetRecommendedDosage(recommendedDosage string)

	// GetItem builds and returns the health product based on the set attributes.
	// It takes a pointer to a HealthBuilder instance as a parameter.
	// It returns the built Health instance.
	GetItem() Health
}

// getBuilder returns a HealthBuilder based on the builderType.
// It takes the builderType as a parameter.
// It returns a HealthBuilder instance or nil if the builderType is not recognized.
func getBuilder(builderType string) HealthBuilder {
	if builderType == "supplements" {
		return newSupplementBuilder()
	}

	if builderType == "skincare" {
		return newSkinCareBuilder()
	}

	if builderType == "haircare" {
		return newHairCareBuilder()
	}

	if builderType == "personalcare" {
		return newPersonalCareBuilder()
	}

	if builderType == "makeup" {
		return newMakeUpBuilder()
	}

	return nil
}
