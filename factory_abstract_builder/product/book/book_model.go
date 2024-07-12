package book

// Book represents a book product.
type Book struct {
	ID          string  // ID is the unique identifier of the book.
	Name        string  // Name is the title of the book.
	Brand       string  // Brand is the brand or publisher of the book.
	Price       float64 // Price is the price of the book.
	Description string  // Description is the brief description of the book.
	Vendor      string  // Vendor is the vendor of the book.
	// category specific attributes
	Author      string // Author is the name of the book's author.
	Publisher   string // Publisher is the name of the book's publisher.
	ISBN        string // ISBN is the International Standard Book Number of the book.
	Pages       string // Pages is the total number of pages in the book.
	Language    string // Language is the language in which the book is written.
	Weight      string // Weight is the weight of the book.
	Dimensions  string // Dimensions are the dimensions of the book.
	Genre       string // Genre is the genre or category of the book.
	Type        string // Type is the type of book (e.g., hardcover, paperback).
	Material    string // Material is the material used for the book's cover.
	Color       string // Color is the color of the book.
	IntendedUse string // IntendedUse is the intended use of the book.
	Size        string // Size is the size of the book.
	Units       string // Number of units of the product in supply
}
