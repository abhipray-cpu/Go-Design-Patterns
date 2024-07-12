package fashion

// createBagProduct creates a bag product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Fashion product.
func createBagProduct(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string) Fashion {
	builder := getBuilder("bags")
	director := NewDirector(builder)
	build := director.buildBag(id, name, price, description, vendor, brand, category, material, size, color, gender, style, prodType)
	return build.GetItem()
}

// createJewelryProduct creates a jewelry product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Fashion product.
func createJewelryProduct(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string, gemStone string) Fashion {
	builder := getBuilder("jewelry")
	director := NewDirector(builder)
	build := director.buildJewelry(id, name, price, description, vendor, brand, category, material, size, color, gender, style, prodType, gemStone)
	return build.GetItem()
}

// createShoeProduct creates a shoe product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Fashion product.
func createShoeProduct(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string) Fashion {
	builder := getBuilder("shoes")
	director := NewDirector(builder)
	build := director.buildShoes(id, name, price, description, vendor, brand, category, material, size, color, gender, style, prodType)
	return build.GetItem()
}

// createClothingProduct creates a clothing product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Fashion product.
func createClothingProduct(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string) Fashion {
	builder := getBuilder("clothing")
	director := NewDirector(builder)
	build := director.buildClothing(id, name, price, description, vendor, brand, category, material, size, color, gender, style, prodType)
	return build.GetItem()
}

// createWatchProduct creates a watch product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Fashion product.
func createWatchProduct(id, name string, price float64, description, vendor string, brand, category string, material string, size string, color string, gender, style, prodType string, resistance string) Fashion {
	builder := getBuilder("watch")
	director := NewDirector(builder)
	build := director.buildWatch(id, name, price, description, vendor, brand, category, material, size, color, gender, style, prodType, resistance)
	return build.GetItem()
}
