package book

import (
	"testing"
)

func TestCreateBookProduct(t *testing.T) {
	book := createBookProduct("123", "Sample Book", 9.99, "Sample description", "Sample Vendor", "Sample Author", "Sample Publisher", "1234567890", "200", "English", "1.2 kg", "20x15x3 cm", "Fiction")

	// Add assertions here to verify the book object
	if book.ID != "123" {
		t.Errorf("Expected ID to be '123', but got '%s'", book.ID)
	}
	if book.Name != "Sample Book" {
		t.Errorf("Expected Name to be 'Sample Book', but got '%s'", book.Name)
	}
	if book.Price != 9.99 {
		t.Errorf("Expected Price to be 9.99, but got %.2f", book.Price)
	}
	if book.Description != "Sample description" {
		t.Errorf("Expected Description to be 'Sample description', but got '%s'", book.Description)
	}
	if book.Vendor != "Sample Vendor" {
		t.Errorf("Expected Vendor to be 'Sample Vendor', but got '%s'", book.Vendor)
	}
	if book.Author != "Sample Author" {
		t.Errorf("Expected Author to be 'Sample Author', but got '%s'", book.Author)
	}
	if book.Publisher != "Sample Publisher" {
		t.Errorf("Expected Publisher to be 'Sample Publisher', but got '%s'", book.Publisher)
	}
	if book.ISBN != "1234567890" {
		t.Errorf("Expected ISBN to be '1234567890', but got '%s'", book.ISBN)
	}
	if book.Pages != "200" {
		t.Errorf("Expected Pages to be '200', but got '%s'", book.Pages)
	}
	if book.Language != "English" {
		t.Errorf("Expected Language to be 'English', but got '%s'", book.Language)
	}
	if book.Weight != "1.2 kg" {
		t.Errorf("Expected Weight to be '1.2 kg', but got '%s'", book.Weight)
	}
	if book.Dimensions != "20x15x3 cm" {
		t.Errorf("Expected Dimensions to be '20x15x3 cm', but got '%s'", book.Dimensions)
	}
	if book.Genre != "Fiction" {
		t.Errorf("Expected Genre to be 'Fiction', but got '%s'", book.Genre)
	}
}

func TestCreateStationaryProduct(t *testing.T) {
	book := CreateStationaryProduct("456", "Sample Stationary Book", 4.99, "Sample description", "Sample Vendor", "0.5 kg", "10x10x1 cm", "Paper", "White", "Writing", "Stationary")

	// Add assertions here to verify the book object
	if book.ID != "456" {
		t.Errorf("Expected ID to be '456', but got '%s'", book.ID)
	}
	if book.Name != "Sample Stationary Book" {
		t.Errorf("Expected Name to be 'Sample Stationary Book', but got '%s'", book.Name)
	}
	if book.Price != 4.99 {
		t.Errorf("Expected Price to be 4.99, but got %.2f", book.Price)
	}
	if book.Description != "Sample description" {
		t.Errorf("Expected Description to be 'Sample description', but got '%s'", book.Description)
	}
	if book.Vendor != "Sample Vendor" {
		t.Errorf("Expected Vendor to be 'Sample Vendor', but got '%s'", book.Vendor)
	}
	if book.Weight != "0.5 kg" {
		t.Errorf("Expected Weight to be '0.5 kg', but got '%s'", book.Weight)
	}
	if book.Dimensions != "10x10x1 cm" {
		t.Errorf("Expected Dimensions to be '10x10x1 cm', but got '%s'", book.Dimensions)
	}
	if book.Material != "Paper" {
		t.Errorf("Expected Material to be 'Paper', but got '%s'", book.Material)
	}
	if book.Color != "White" {
		t.Errorf("Expected Color to be 'White', but got '%s'", book.Color)
	}
	if book.IntendedUse != "Writing" {
		t.Errorf("Expected Intended to be 'Writing', but got '%s'", book.IntendedUse)
	}
	if book.Type != "Stationary" {
		t.Errorf("Expected Type to be 'Stationary', but got '%s'", book.Type)
	}
}

func TestCreateArtCraftProduct(t *testing.T) {
	book := CreateArtCraftProduct("789", "Sample ArtCraft Book", 19.99, "Sample description", "Sample Vendor", "1.5 kg", "30x20x5 cm", "Canvas", "Multicolor", "Painting", "ArtCraft")

	// Add assertions here to verify the book object
	if book.ID != "789" {
		t.Errorf("Expected ID to be '789', but got '%s'", book.ID)
	}
	if book.Name != "Sample ArtCraft Book" {
		t.Errorf("Expected Name to be 'Sample ArtCraft Book', but got '%s'", book.Name)
	}
	if book.Price != 19.99 {
		t.Errorf("Expected Price to be 19.99, but got %.2f", book.Price)
	}
	if book.Description != "Sample description" {
		t.Errorf("Expected Description to be 'Sample description', but got '%s'", book.Description)
	}
	if book.Vendor != "Sample Vendor" {
		t.Errorf("Expected Vendor to be 'Sample Vendor', but got '%s'", book.Vendor)
	}
	if book.Weight != "1.5 kg" {
		t.Errorf("Expected Weight to be '1.5 kg', but got '%s'", book.Weight)
	}
	if book.Dimensions != "30x20x5 cm" {
		t.Errorf("Expected Dimensions to be '30x20x5 cm', but got '%s'", book.Dimensions)
	}
	if book.Material != "Canvas" {
		t.Errorf("Expected Material to be 'Canvas', but got '%s'", book.Material)
	}
	if book.Color != "Multicolor" {
		t.Errorf("Expected Color to be 'Multicolor', but got '%s'", book.Color)
	}
	if book.IntendedUse != "Painting" {
		t.Errorf("Expected Intended to be 'Painting', but got '%s'", book.IntendedUse)
	}
	if book.Type != "ArtCraft" {
		t.Errorf("Expected Type to be 'ArtCraft', but got '%s'", book.Type)
	}
}

func TestCreateSupplyProduct(t *testing.T) {
	book := CreateSupplyProduct("987", "Sample Supply Book", 2.99, "Sample description", "Sample Vendor", "0.2 kg", "5x5x1 cm", "Studying", "Supplies", "10")

	// Add assertions here to verify the book object
	if book.ID != "987" {
		t.Errorf("Expected ID to be '987', but got '%s'", book.ID)
	}
	if book.Name != "Sample Supply Book" {
		t.Errorf("Expected Name to be 'Sample Supply Book', but got '%s'", book.Name)
	}
	if book.Price != 2.99 {
		t.Errorf("Expected Price to be 2.99, but got %.2f", book.Price)
	}
	if book.Description != "Sample description" {
		t.Errorf("Expected Description to be 'Sample description', but got '%s'", book.Description)
	}
	if book.Vendor != "Sample Vendor" {
		t.Errorf("Expected Vendor to be 'Sample Vendor', but got '%s'", book.Vendor)
	}
	if book.Weight != "0.2 kg" {
		t.Errorf("Expected Weight to be '0.2 kg', but got '%s'", book.Weight)
	}
	if book.Dimensions != "5x5x1 cm" {
		t.Errorf("Expected Dimensions to be '5x5x1 cm', but got '%s'", book.Dimensions)
	}
	if book.IntendedUse != "Studying" {
		t.Errorf("Expected Intended to be 'Studying', but got '%s'", book.IntendedUse)
	}
	if book.Type != "Supplies" {
		t.Errorf("Expected Type to be 'Supplies', but got '%s'", book.Type)
	}
	if book.Units != "10" {
		t.Errorf("Expected Units to be '10', but got '%s'", book.Units)
	}
}

// Add assertions here to verify the book object
