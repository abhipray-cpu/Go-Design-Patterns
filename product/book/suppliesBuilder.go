package book

// SuppliesBuilder is a struct that represents a builder for creating supplies.
type SuppliesBuilder struct {
	book Book
}

// newSupplyBuilder creates a new instance of SuppliesBuilder.
func newSupplyBuilder() *SuppliesBuilder {
	return &SuppliesBuilder{}
}

// SetAttributes sets the attributes of the SuppliesBuilder's book.
// It takes the following parameters:
// - id: The ID of the book.
// - name: The name of the book.
// - price: The price of the book.
// - description: The description of the book.
// - vendor: The vendor of the book.
func (b *SuppliesBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	b.book.ID = id
	b.book.Name = name
	b.book.Price = price
	b.book.Description = description
	b.book.Vendor = vendor
}

// SetWeight sets the weight of the book.
func (b *SuppliesBuilder) SetWeight(weight string) {
	b.book.Weight = weight
}

// SetDimensions sets the dimensions of the book.
func (b *SuppliesBuilder) SetDimensions(dimensions string) {
	b.book.Dimensions = dimensions
}

// SetUnits sets the units of the book.
func (b *SuppliesBuilder) SetUnits(units string) {
	b.book.Units = units
}

// SetIntendedUse sets the intended use of the book.
func (b *SuppliesBuilder) SetIntendedUse(intendedUse string) {
	b.book.IntendedUse = intendedUse
}

// SetType sets the type of the book.
func (b *SuppliesBuilder) SetType(bookType string) {
	b.book.Type = bookType
}

// GetItem returns the built book.
func (b *SuppliesBuilder) GetItem() Book {
	return b.book
}

// unused methods from abstract builder interface

func (b *SuppliesBuilder) SetMaterial(material string) {
}

func (b *SuppliesBuilder) SetSize(size string) {
}

func (b *SuppliesBuilder) SetGenre(genre string) {
}

func (b *SuppliesBuilder) SetAuthor(author string) {
}

func (b *SuppliesBuilder) SetPublisher(publisher string) {
}

func (b *SuppliesBuilder) SetISBN(isbn string) {
}

func (b *SuppliesBuilder) SetColor(color string) {
}

func (b *SuppliesBuilder) SetLanguage(language string) {
}

func (b *SuppliesBuilder) SetPages(pages string) {
}