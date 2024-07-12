package pet

// createAccessory creates an accessory product for pets.
// It takes in various parameters such as id, name, price, description, vendor, category, brand, suitable, durability, weight, material, size, and prodType.
// It uses the builder pattern to construct the accessory product and returns the built product.
func createAccessory(id, name string, price float64, description, vendor string, category string, brand string, suitable []string, durability string, weight, material, size, prodType string) Pet {
	builder := getBuilder("accessory")
	director := newDirector(builder)
	build := director.buildAccessoryProduct(id, name, price, description, vendor, category, brand, suitable, durability, weight, material, size, prodType)
	return build.GetItem()
}

// createFood creates a food product for pets.
// It takes in various parameters such as id, name, price, description, vendor, category, brand, ingredients, nutritionalValue, weight, size, prodType, and suitable.
// It uses the builder pattern to construct the food product and returns the built product.
func createFood(id, name string, price float64, description, vendor string, category string, brand string, ingredients []string, nutritionalValue string, weight, size, prodType string, suitable []string) Pet {
	builder := getBuilder("food")
	director := newDirector(builder)
	build := director.buildFoodProduct(id, name, price, description, vendor, category, brand, ingredients, nutritionalValue, weight, size, prodType, suitable)
	return build.GetItem()
}

// createGrooming creates a grooming product for pets.
// It takes in various parameters such as id, name, price, description, vendor, category, brand, suitable, durability, volume, size, prodType, and ingredients.
// It uses the builder pattern to construct the grooming product and returns the built product.
func createGrooming(id, name string, price float64, description, vendor string, category string, brand string, suitable []string, durability string,
	volume string, size string, prodType string, ingredients []string) Pet {
	builder := getBuilder("grooming")
	director := newDirector(builder)
	build := director.buildGroomingProduct(id, name, price, description, vendor, category, brand, suitable, durability,
		volume, size, prodType, ingredients)
	return build.GetItem()
}

// createToy creates a toy product for pets.
// It takes in various parameters such as id, name, price, description, vendor, category, brand, suitable, durability, material, prodType, and size.
// It uses the builder pattern to construct the toy product and returns the built product.
func createToy(id, name string, price float64, description, vendor string, category string, brand string, suitable []string, durability string, material, prodType, size string) Pet {
	builder := getBuilder("toy")
	director := newDirector(builder)
	build := director.buildToyProduct(id, name, price, description, vendor, category, brand, suitable, durability, material, prodType, size)
	return build.GetItem()
}
