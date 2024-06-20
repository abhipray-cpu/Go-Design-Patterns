package health

// createHairCareProduct creates a hair care product using the specified parameters.
// It uses the hair care product builder to construct the product.
// Returns the created health product.
func createHairCareProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, instructions []string, prodType string, spf string, volume string, sulfate bool, gender string) Health {
	builder := getBuilder("haircare")
	director := NewDirector(builder)
	build := director.buildHairCareProduct(id, name, price, description, vendor, category, ingredients, brand, instructions, prodType, spf, volume, sulfate, gender)
	return build.GetItem()
}

// createMakeupProduct creates a makeup product using the specified parameters.
// It uses the makeup product builder to construct the product.
// Returns the created health product.
func createMakeupProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, prodType string, volume, color, finish string) Health {
	builder := getBuilder("makeup")
	director := NewDirector(builder)
	build := director.builMakeupProduct(id, name, price, description, vendor, category, ingredients, brand, prodType, volume, color, finish)
	return build.GetItem()
}

// createPersonalCareProduct creates a personal care product using the specified parameters.
// It uses the personal care product builder to construct the product.
// Returns the created health product.
func createPersonalCareProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, instructions []string, prodType string, volume string, gender string) Health {
	builder := getBuilder("personalcare")
	director := NewDirector(builder)
	build := director.buildPersonalCareProduct(id, name, price, description, vendor, category, ingredients, brand, instructions, prodType, volume, gender)
	return build.GetItem()
}

// createSkinCareProduct creates a skin care product using the specified parameters.
// It uses the skin care product builder to construct the product.
// Returns the created health product.
func createSkinCareProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, instructions []string, prodType string, spf string, volume string, gender string) Health {
	builder := getBuilder("skincare")
	director := NewDirector(builder)
	build := director.buildSkinCareProduct(id, name, price, description, vendor, category, ingredients, brand, instructions, prodType, spf, volume, gender)
	return build.GetItem()
}

// createSupplementProduct creates a supplement product using the specified parameters.
// It uses the supplement product builder to construct the product.
// Returns the created health product.
func createSupplementProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, prodType string, volume string, form string, dosage string, gender string, instructions []string) Health {
	builder := getBuilder("supplements")
	director := NewDirector(builder)
	build := director.buildSupplementProduct(id, name, price, description, vendor, category, ingredients, brand, prodType, volume, form, dosage, gender, instructions)
	return build.GetItem()
}
