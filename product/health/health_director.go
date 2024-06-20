package health

type Director struct {
	builder HealthBuilder
}

// NewDirector creates a new Director instance with the given HealthBuilder.
func NewDirector(b HealthBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// setBuilder sets the builder for the Director.
// The builder is responsible for constructing the Health object.
func (d *Director) setBuilder(b HealthBuilder) {
	d.builder = b
}

// buildHairCareProduct builds a hair care product using the provided attributes and returns a HealthBuilder.
// It sets the attributes such as ID, name, price, description, vendor, category, ingredients, brand, instructions,
// product type, SPF, volume, sulfate, and gender on the builder.
func (d *Director) buildHairCareProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, instructions []string, prodType string, spf string, volume string, sulfate bool, gender string) HealthBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetIngredients(ingredients)
	d.builder.SetBrand(brand)
	d.builder.SetUsageInstructions(instructions)
	d.builder.SetType(prodType)
	d.builder.SetSPF(spf)
	d.builder.SetVolume(volume)
	d.builder.SetSulfateFree(sulfate)
	d.builder.SetForMenWomen(gender)
	return d.builder
}

// buildSkinCareProduct builds a skin care product using the provided attributes and returns a HealthBuilder.
// The skin care product is built by setting the ID, name, price, description, and vendor using the provided values.
// The category, ingredients, brand, usage instructions, product type, SPF, and volume are also set using the provided values.
func (d *Director) buildSkinCareProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, instructions []string, prodType string, spf string, volume string, gender string) HealthBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetIngredients(ingredients)
	d.builder.SetBrand(brand)
	d.builder.SetUsageInstructions(instructions)
	d.builder.SetType(prodType)
	d.builder.SetSPF(spf)
	d.builder.SetVolume(volume)
	d.builder.SetForMenWomen(gender)
	return d.builder
}

// buildSupplementProduct builds a supplement product using the provided attributes and returns the HealthBuilder instance.
func (d *Director) buildSupplementProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, prodType string, volume string, form string, dosage string, gender string, instructions []string) HealthBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetIngredients(ingredients)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetVolume(volume)
	d.builder.SetSupplementForm(form)
	d.builder.SetRecommendedDosage(dosage)
	d.builder.SetForMenWomen(gender)
	d.builder.SetUsageInstructions(instructions)
	return d.builder
}

// buildPersonalCareProduct builds a personal care product using the provided attributes and returns a HealthBuilder.
// It sets the attributes such as id, name, price, description, and vendor using the builder's SetAttributes method.
// It also sets the category, ingredients, brand, usage instructions, product type, volume, and gender using the respective builder methods.
func (d *Director) buildPersonalCareProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, instructions []string, prodType string, volume string, gender string) HealthBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetIngredients(ingredients)
	d.builder.SetBrand(brand)
	d.builder.SetUsageInstructions(instructions)
	d.builder.SetType(prodType)
	d.builder.SetVolume(volume)
	d.builder.SetForMenWomen(gender)
	return d.builder
}

// builMakeupProduct builds a makeup product using the provided attributes and returns a HealthBuilder instance.
// It sets the attributes such as id, name, price, description, and vendor using the builder's SetAttributes method.
// It also sets the category, ingredients, brand, product type, volume, color, and finish using the respective builder methods.
func (d *Director) builMakeupProduct(id, name string, price float64, description, vendor string, category string, ingredients []string, brand string, prodType string, volume, color, finish string) HealthBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetCategory(category)
	d.builder.SetIngredients(ingredients)
	d.builder.SetBrand(brand)
	d.builder.SetType(prodType)
	d.builder.SetVolume(volume)
	d.builder.SetColor(color)
	d.builder.SetFinish(finish)
	return d.builder
}
