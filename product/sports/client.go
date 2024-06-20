package sports

// createCampingProduct creates a camping product using the specified parameters.
// It uses the builder pattern to construct the product.
func createCampingProduct(id, name string, price float64, description string, category string, minAge, maxAge int, dimension string, durability string, material string, weight string, vendor string, prodType, size, brand, color string, suitable []string) Sports {
	builder := getBuilder("Camping")
	director := newDirector(builder)
	build := director.buildSportsProduct(id, name, price, description, category, minAge, maxAge, dimension, durability, material, weight, vendor, prodType, size, brand, color, suitable)
	return build.GetItem()
}

// createEquipmentProduct creates an equipment product using the specified parameters.
// It uses the builder pattern to construct the product.
func createEquipmentProduct(id, name string, price float64, description string, category string, minAge, maxAge int, dimension string, durability string, material string, weight string, vendor string, prodType, size, brand, color string, suitable []string) Sports {
	builder := getBuilder("Equipment")
	director := newDirector(builder)
	build := director.buildSportsProduct(id, name, price, description, category, minAge, maxAge, dimension, durability, material, weight, vendor, prodType, size, brand, color, suitable)
	return build.GetItem()
}

// createExerciseProduct creates an exercise product using the specified parameters.
// It uses the builder pattern to construct the product.
func createExerciseProduct(id, name string, price float64, description string, category string, minAge, maxAge int, dimension string, durability string, material string, weight string, vendor string, prodType, size, brand, color string, suitable []string) Sports {
	builder := getBuilder("Exercise")
	director := newDirector(builder)
	build := director.buildSportsProduct(id, name, price, description, category, minAge, maxAge, dimension, durability, material, weight, vendor, prodType, size, brand, color, suitable)
	return build.GetItem()
}

// createRecreationProduct creates a recreation product using the specified parameters.
// It uses the builder pattern to construct the product.
func createRecreationProduct(id, name string, price float64, description string, category string, minAge, maxAge int, dimension string, durability string, material string, weight string, vendor string, prodType, size, brand, color string, suitable []string) Sports {
	builder := getBuilder("Recreation")
	director := newDirector(builder)
	build := director.buildSportsProduct(id, name, price, description, category, minAge, maxAge, dimension, durability, material, weight, vendor, prodType, size, brand, color, suitable)
	return build.GetItem()
}
