package book

// Director represents a director that oversees the construction of a book.
type Director struct {
	builder BookBuilder
}

// NewDirector creates a new Director instance with the given BookBuilder.
func NewDirector(builder BookBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// SetBuilder sets the builder for the Director.
// It takes a BookBuilder as a parameter and assigns it to the builder field of the Director.
func (d *Director) SetBuilder(builder BookBuilder) {
	d.builder = builder
}

// buildBook is a method of the Director struct that builds a book using the provided attributes and returns a BookBuilder.
// It sets the common attributes of the book such as id, name, price, description, and vendor.
// It also sets the specific attributes of the book such as author, publisher, ISBN, pages, language, weight, dimensions, and genre.
func (d *Director) buildBook(id, name string, price float64, description, vendor string, author string, publisher string, isbn string, pages string, language string, weight, dimension, genre string) BookBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetAuthor(author)
	d.builder.SetPublisher(publisher)
	d.builder.SetISBN(isbn)
	d.builder.SetPages(pages)
	d.builder.SetLanguage(language)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetGenre(genre)
	return d.builder
}

// buildStationary builds a stationary book using the provided attributes and returns the BookBuilder.
// It sets the attributes of the book such as id, name, price, description, and vendor.
// It also sets additional attributes specific to stationary books such as weight, dimension, material, color, intended use, and product type.
func (d *Director) buildStationary(id, name string, price float64, description, vendor string, weight, dimension, material, color string, use string, prodType string) BookBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetMaterial(material)
	d.builder.SetColor(color)
	d.builder.SetIntendedUse(use)
	d.builder.SetType(prodType)
	return d.builder
}

// buildArtCraft builds a book using the provided attributes and returns a BookBuilder.
// It sets the attributes of the book such as id, name, price, description, and vendor.
// It also sets additional attributes such as weight, dimensions, material, color, intended use, and product type.
// The BookBuilder instance is then returned.
func (d *Director) buildArtCraft(id, name string, price float64, description, vendor string, weight, dimension, material, color string, use string, prodType string) BookBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetMaterial(material)
	d.builder.SetColor(color)
	d.builder.SetIntendedUse(use)
	d.builder.SetType(prodType)
	return d.builder
}

// buildSupply builds a Book using the provided attributes and returns the BookBuilder.
// It sets the attributes of the Book using the given id, name, price, description, and vendor.
// It also sets the weight, dimensions, intended use, product type, and units of the Book.
func (d *Director) buildSupply(id, name string, price float64, description, vendor string, weight, dimension string, use string, prodType string, units string) BookBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetIntendedUse(use)
	d.builder.SetType(prodType)
	d.builder.SetUnits(units)
	return d.builder
}
