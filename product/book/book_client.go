package book

// createBookProduct creates a new Book product with the specified attributes.
// It uses the Builder pattern to construct the Book object.
// Parameters:
//   - id: The ID of the book.
//   - name: The name of the book.
//   - price: The price of the book.
//   - description: The description of the book.
//   - vendor: The vendor of the book.
//   - author: The author of the book.
//   - publisher: The publisher of the book.
//   - isbn: The ISBN of the book.
//   - pages: The number of pages in the book.
//   - language: The language of the book.
//   - weight: The weight of the book.
//   - dimension: The dimensions of the book.
//   - genre: The genre of the book.
// Returns:
//   The created Book object.
func createBookProduct(id, name string, price float64, description, vendor string, author string, publisher string, isbn string, pages string, language string, weight, dimension, genre string) Book {
	builder := getBookBuilder("book")
	director := NewDirector(builder)
	build := director.buildBook(id, name, price, description, vendor, author, publisher, isbn, pages, language, weight, dimension, genre)
	return build.GetItem()
}

// CreateStationaryProduct creates a stationary book product with the specified attributes.
// It takes in the ID, name, price, description, vendor, weight, dimension, material, color, use, and prodType of the book.
// It returns a Book object representing the created stationary book product.
func CreateStationaryProduct(id, name string, price float64, description, vendor string, weight, dimension, material, color string, use string, prodType string) Book {
	builder := getBookBuilder("stationary")
	director := NewDirector(builder)
	build := director.buildStationary(id, name, price, description, vendor, weight, dimension, material, color, use, prodType)
	return build.GetItem()
}

// CreateArtCraftProduct creates a new artcraft book product with the specified attributes.
// It takes in the ID, name, price, description, vendor, weight, dimension, material, color, use, and prodType of the book.
// It returns a Book object representing the created artcraft book.
func CreateArtCraftProduct(id, name string, price float64, description, vendor string, weight, dimension, material, color string, use string, prodType string) Book {
	builder := getBookBuilder("artcraft")
	director := NewDirector(builder)
	build := director.buildArtCraft(id, name, price, description, vendor, weight, dimension, material, color, use, prodType)
	return build.GetItem()
}

// CreateSupplyProduct creates a supply product of type Book with the given parameters.
// It uses a builder pattern to construct the product and returns the created Book.
func CreateSupplyProduct(id, name string, price float64, description, vendor string, weight, dimension string, use string, prodType string, units string) Book {
	builder := getBookBuilder("supplies")
	director := NewDirector(builder)
	build := director.buildSupply(id, name, price, description, vendor, weight, dimension, use, prodType, units)
	return build.GetItem()
}
