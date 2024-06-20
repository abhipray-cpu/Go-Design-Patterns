package home

// createBathProduct creates a bath product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Home object.
func createBathProduct(id, name string, price float64, description, vendor string, category string, brand, prodType string, volume, scent, skinType string, material, dimension, weight, color, capacity, stye string, assembly bool) Home {
	builder := getBuilder("bath")
	director := NewDirector(builder)
	build := director.buildBath(id, name, price, description, vendor, category, brand, prodType, volume, scent, skinType, material, dimension, weight, color, capacity, stye, assembly)
	return build.GetItem()
}

// createDecorProduct creates a decor product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Home object.
func createDecorProduct(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimension string, assembly bool) Home {
	builder := getBuilder("decor")
	director := NewDirector(builder)
	build := director.buildDecor(id, name, price, description, vendor, category, brand, prodType, color, material, weight, style, dimension, assembly)
	return build.GetItem()
}

// createFurnitureProduct creates a furniture product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Home object.
func createFurnitureProduct(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimension string, assembly bool, capacity string) Home {
	builder := getBuilder("furniture")
	director := NewDirector(builder)
	build := director.buildFurniture(id, name, price, description, vendor, category, brand, prodType, color, material, weight, style, dimension, assembly, capacity)
	return build.GetItem()
}

// createKitchenProduct creates a kitchen product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Home object.
func createKitchenProduct(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimension, capacity string, safe bool, assembly bool, power string) Home {
	builder := getBuilder("kitchen")
	director := NewDirector(builder)
	build := director.buildKitchen(id, name, price, description, vendor, category, brand, prodType, color, material, weight, style, dimension, capacity, safe, assembly, power)
	return build.GetItem()
}

// createToolProduct creates a tool product using the specified parameters.
// It uses the builder pattern to construct the product.
// Returns the created Home object.
func createToolProduct(id, name string, price float64, description, vendor string, category string, brand, prodType string, color, material, weight, style, dimensions, powerSource string) Home {
	builder := getBuilder("tools")
	director := NewDirector(builder)
	build := director.buildTools(id, name, price, description, vendor, category, brand, prodType, color, material, weight, style, dimensions, powerSource)
	return build.GetItem()
}
