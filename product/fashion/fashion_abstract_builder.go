package fashion

// FashionBuilder is a struct that represents a builder for creating Fashion objects.
// FashionBuilder is an interface that defines the methods for building a fashion product.
type FashionBuilder interface {
	// SetAttributes sets the attributes of the fashion product.
	// It takes the ID, name, price, description, and vendor as parameters.
	SetAttributes(id, name string, price float64, description, vendor string)

	// SetBrand sets the brand of the fashion product.
	// It takes the brand name as a parameter.
	SetBrand(brand string)

	// SetMaterial sets the material of the fashion product.
	// It takes the material name as a parameter.
	SetMaterial(material string)

	// SetCategory sets the category of the fashion product.
	// It takes the category name as a parameter.
	SetCategory(category string)

	// SetSize sets the size of the fashion product.
	// It takes the size name as a parameter.
	SetSize(size string)

	// SetColor sets the color of the fashion product.
	// It takes the color name as a parameter.
	SetColor(color string)

	// SetGender sets the gender of the fashion product.
	// It takes the gender name as a parameter.
	SetGender(gender string)

	// SetStyle sets the style of the fashion product.
	// It takes the style name as a parameter.
	SetStyle(style string)

	// SetType sets the type of the fashion product.
	// It takes the fashion type name as a parameter.
	SetType(fashionType string)

	// SetGemStone sets the gemstone of the fashion product.
	// It takes the gemstone name as a parameter.
	SetGemStone(gemStone string)

	// SetWaterResistance sets the water resistance of the fashion product.
	// It takes the water resistance value as a parameter.
	SetWaterResistance(waterResistance string)

	// GetFashion returns the built fashion product.
	GetItem() Fashion
}

// getBuilder returns a FashionBuilder based on the given builderType.
func getBuilder(builderType string) FashionBuilder {
	if builderType == "bags" {
		return newBagBuilder()
	}
	if builderType == "jewelry" {
		return newJewelryBuilder()
	}
	if builderType == "shoes" {
		return newShoesBuilder()
	}
	if builderType == "watch" {
		return newWatchBuilder()
	}
	if builderType == "clothing" {
		return newClothingBuilder()
	}
	return nil
}
