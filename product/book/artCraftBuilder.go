package book

// ArtCraftBuilder is a struct that represents a builder for creating art and craft books.
type ArtCraftBuilder struct {
	book Book
}

// newArtCraftBuilder creates a new instance of ArtCraftBuilder.
func newArtCraftBuilder() *ArtCraftBuilder {
	return &ArtCraftBuilder{}
}

// SetAttributes sets the attributes of the ArtCraftBuilder's book.
// It takes the following parameters:
// - id: The ID of the book.
// - name: The name of the book.
// - price: The price of the book.
// - description: The description of the book.
// - vendor: The vendor of the book.
func (b *ArtCraftBuilder) SetAttributes(id, name string, price float64, description string, vendor string) {
	b.book.ID = id
	b.book.Name = name
	b.book.Price = price
	b.book.Description = description
	b.book.Vendor = vendor
}

// SetWeight sets the weight of the book.
func (b *ArtCraftBuilder) SetWeight(weight string) {
	b.book.Weight = weight
}

// SetDimensions sets the dimensions of the book.
func (b *ArtCraftBuilder) SetDimensions(dimensions string) {
	b.book.Dimensions = dimensions
}

// SetMaterial sets the material of the book.   
func (b *ArtCraftBuilder) SetMaterial(material string) {
	b.book.Material = material
}

// SetColor sets the color of the book.
func (b *ArtCraftBuilder) SetColor(color string) {
	b.book.Color = color
}

// SetIntendedUse sets the intended use of the book.
func (b *ArtCraftBuilder) SetIntendedUse(intendedUse string) {
	b.book.IntendedUse = intendedUse
}

// SetType sets the type of the book.
func (b *ArtCraftBuilder) SetType(bookType string) {
	b.book.Type = bookType
}


// SetSize sets the size of the book.
func (b *ArtCraftBuilder) SetSize(size string) {
	b.book.Size = size
}

// GetItem returns the book created by the ArtCraftBuilder.
func (b *ArtCraftBuilder) GetItem() Book {
	return b.book
}

// unused methods from abstract builder interface
func (b *ArtCraftBuilder) SetUnits(units string) {
}

func (b *ArtCraftBuilder) SetGenre(genre string) {

}

func (b *ArtCraftBuilder) SetAuthor(author string) {

}

func (b *ArtCraftBuilder) SetPublisher(publisher string) {
}

func (b *ArtCraftBuilder) SetISBN(isbn string) {
}

func (b *ArtCraftBuilder) SetPages(pages string) {
}

func (b *ArtCraftBuilder) SetLanguage(language string) {
}
