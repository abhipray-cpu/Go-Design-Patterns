package book

// StationaryBuilder is a builder for creating stationary books.
type StationaryBuilder struct {
	book Book
}

// newStationaryBuilder creates a new instance of StationaryBuilder.
func newStationaryBuilder() *StationaryBuilder {
	return &StationaryBuilder{}
}

// SetAttributes sets the attributes of the book.
func (b *StationaryBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	b.book.ID = id
	b.book.Name = name
	b.book.Price = price
	b.book.Description = description
	b.book.Vendor = vendor
}

// SetWeight sets the weight of the book.
func (b *StationaryBuilder) SetWeight(weight string) {
	b.book.Weight = weight
}

// SetDimensions sets the dimensions of the book.
func (b *StationaryBuilder) SetDimensions(dimensions string) {
	b.book.Dimensions = dimensions
}

// SetMaterial sets the material of the book.
func (b *StationaryBuilder) SetMaterial(material string) {
	b.book.Material = material
}

// SetColor sets the color of the book.
func (b *StationaryBuilder) SetColor(color string) {
	b.book.Color = color
}

// SetIntendedUse sets the intended use of the book.
func (b *StationaryBuilder) SetIntendedUse(intendedUse string) {
	b.book.IntendedUse = intendedUse
}

// SetType sets the type of the book.
func (b *StationaryBuilder) SetType(bookType string) {
	b.book.Type = bookType
}

// GetItem returns the built book.
func (b *StationaryBuilder) GetItem() Book {
	return b.book
}

// unused functions from the abstract interface

// SetUnits sets the units of the book.
func (b *StationaryBuilder) SetUnits(units string) {
}

// SetSize sets the size of the book.
func (b *StationaryBuilder) SetSize(size string) {

}

// SetGenre sets the genre of the book.
func (b *StationaryBuilder) SetGenre(genre string) {

}

// SetAuthor sets the author of the book.
func (b *StationaryBuilder) SetAuthor(author string) {

}

// SetPublisher sets the publisher of the book.
func (b *StationaryBuilder) SetPublisher(publisher string) {
}

// SetISBN sets the ISBN of the book.
func (b *StationaryBuilder) SetISBN(isbn string) {
}

// SetPages sets the number of pages of the book.
func (b *StationaryBuilder) SetPages(pages string) {
}

// SetLanguage sets the language of the book.
func (b *StationaryBuilder) SetLanguage(language string) {
}