package book

// BookBuilder is an interface that defines the methods for building a book.
type BookBuilder interface {
	// SetAttributes sets the attributes of the book.
	// It takes the ID, name, price, and description of the book as parameters.
	SetAttributes(id, name string, price float64, description string, vendor string)

	// SetAuthor sets the author of the book.
	// It takes the author name as a parameter.
	SetAuthor(author string)

	// SetPublisher sets the publisher of the book.
	// It takes the publisher name as a parameter.
	SetPublisher(publisher string)

	// SetISBN sets the ISBN of the book.
	// It takes the ISBN as a parameter.
	SetISBN(isbn string)

	// SetPages sets the number of pages of the book.
	// It takes the number of pages as a parameter.
	SetPages(pages string)

	// SetLanguage sets the language of the book.
	// It takes the language as a parameter.
	SetLanguage(language string)

	// SetWeight sets the weight of the book.
	// It takes the weight as a parameter.
	SetWeight(weight string)

	// SetDimensions sets the dimensions of the book.
	// It takes the dimensions as a parameter.
	SetDimensions(dimensions string)

	// SetGenre sets the genre of the book.
	// It takes the genre as a parameter.
	SetGenre(genre string)

	// SetType sets the type of the book.
	// It takes the book type as a parameter.
	SetType(bookType string)

	// SetMaterial sets the material of the book.
	// It takes the material as a parameter.
	SetMaterial(material string)

	// SetColor sets the color of the book.
	// It takes the color as a parameter.
	SetColor(color string)

	// SetIntendedUse sets the intended use of the book.
	// It takes the intended use as a parameter.
	SetIntendedUse(intendedUse string)

	// SetSize sets the size of the book.
	// It takes the size as a parameter.
	SetSize(size string)

	// SetUnits sets the units of measurement for the book.
	// It takes the units as a parameter.
	SetUnits(units string)

	// GetBook returns the built Book object.
	GetItem() Book
}

// getBookBuilder returns a BookBuilder based on the provided builderType.
// It creates and returns a specific type of builder based on the input string.
// If the builderType is "book", it returns a newBooksBuilder.
// If the builderType is "stationary", it returns a newStationaryBuilder.
// If the builderType is "artcraft", it returns a newArtCraftBuilder.
// If the builderType is "supplies", it returns a newSupplyBuilder.
// If the builderType is not recognized, it returns nil.
func getBookBuilder(builderType string) BookBuilder {
	if builderType == "book" {
		return newBooksBuilder()
	}

	if builderType == "stationary" {
		return newStationaryBuilder()
	}
	if builderType == "artcraft" {
		return newArtCraftBuilder()
	}
	if builderType == "supplies" {
		return newSupplyBuilder()
	}
	return nil
}
