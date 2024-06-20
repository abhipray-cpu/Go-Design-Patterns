package book

// BooksBuilder is a struct that represents a builder for creating books.
type BooksBuilder struct {
	book Book
}

// newBooksBuilder creates a new instance of BooksBuilder.
func newBooksBuilder() *BooksBuilder {
	return &BooksBuilder{}
}

// SetAttributes sets the attributes of the book being built.
// It takes the following parameters:
// - id: The ID of the book.
// - name: The name of the book.
// - price: The price of the book.
// - description: The description of the book.
// - vendor: The vendor of the book.
func (b *BooksBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	b.book.ID = id
	b.book.Name = name
	b.book.Price = price
	b.book.Description = description
	b.book.Vendor = vendor
}

// SetAuthor sets the author of the book.
func (b *BooksBuilder) SetAuthor(author string) {
	b.book.Author = author

}

// SetPublisher sets the publisher of the book.
func (b *BooksBuilder) SetPublisher(publisher string) {
	b.book.Publisher = publisher

}

// SetISBN sets the ISBN of the book.
func (b *BooksBuilder) SetISBN(isbn string) {
	b.book.ISBN = isbn

}

// SetPages sets the number of pages of the book.
func (b *BooksBuilder) SetPages(pages string) {
	b.book.Pages = pages

}

// SetLanguage sets the language of the book.
func (b *BooksBuilder) SetLanguage(language string) {
	b.book.Language = language

}

// SetWeight sets the weight of the book.
func (b *BooksBuilder) SetWeight(weight string) {
	b.book.Weight = weight

}

// SetDimensions sets the dimensions of the book.
func (b *BooksBuilder) SetDimensions(dimensions string) {
	b.book.Dimensions = dimensions

}

// SetGenre sets the genre of the book.
func (b *BooksBuilder) SetGenre(genre string) {
	b.book.Genre = genre

}

// GetItem returns the book created by the BooksBuilder.
func (b *BooksBuilder) GetItem() Book {
	return b.book
}

// unused functions from the abstract interface
func (b *BooksBuilder) SetType(bookType string) {

}

func (b *BooksBuilder) SetMaterial(material string) {

}

func (b *BooksBuilder) SetColor(color string) {

}

func (b *BooksBuilder) SetIntendedUse(intendedUse string) {

}

func (b *BooksBuilder) SetSize(size string) {

}

func (b *BooksBuilder) SetUnits(units string) {

}
