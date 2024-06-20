package toys

// ToyBuilder is an interface that defines the methods for building a toy.
type ToyBuilder interface {
	// this method allows for setting the attributes of the toy
	SetAttributes(id, name string, price float64, description string)
	// this method allows for setting the category of the toy
	SetCategory(category string)
	// this method allows for setting the age range of the toy
	SetAgeRange(minAge, maxAge int)
	// this method allows for setting the brand of the toy
	SetBrand(brand string)
	// this method allows for setting the weight of the toy
	SetWeight(weight string)
	// this method allows for setting the material of the toy
	SetMaterial(material string)
	// this method allows for setting the color of the toy
	SetColor(color string)
	// this method allows for setting the size of the toy
	SetSize(size string)
	// this method allows for setting the type of the toy
	SetType(prodType string)
	// this methods allows for setting dimension of the product
	SetDimensions(dimension string)
	// setting the age group suitable for the toy
	SetSuitableFor(suitableFor []string)
	// this method allows for setting the vendor of the toy
	SetVendor(vendor string)
	// this method allows for getting the toy
	GetToy() Toy
}

// getBuilder returns a ToyBuilder based on the given builderType.
// It accepts a string parameter builderType and returns a ToyBuilder interface.
// If the builderType is "Action", it returns a new instance of ActionToyBuilder.
// If the builderType is "Board", it returns a new instance of BoardToyBuilder.
// If the builderType is "Video", it returns a new instance of VideoToyBuilder.
// If the builderType is not one of the above, it returns nil.
func getBuilder(builderType string) ToyBuilder {
	switch builderType {
	case "Action":
		return NewActionToyBuilder()
	case "Board":
		return NewBoardToyBuilder()
	case "Video":
		return NewVideoToyBuilder()
	default:
		return nil
	}
}
