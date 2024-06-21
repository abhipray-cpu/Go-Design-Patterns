package toys

import (
	"testing"
)

func TestCreateActionToy(t *testing.T) {
	toy := createActionToy("123", "Action Toy", 9.99, "Description", "Category", 3, 8, "Brand", "1kg", "Vendor", "Material", "Color", "Size", "ProdType", "Dimension", []string{"Suitable"})

	// Add assertions here to validate the toy object

	// Example assertion:
	if toy.ID != "123" {
		t.Errorf("Expected ID to be '123', got '%s'", toy.ID)
	}
}

func TestCreateBoardGame(t *testing.T) {
	toy := createBoardGame("456", "Board Game", 19.99, "Description", "Category", 5, 12, "Brand", "2kg", "Vendor", "ProdType", "Dimension", []string{"Suitable"})

	// Add assertions here to validate the toy object

	// Example assertion:
	if toy.ID != "456" {
		t.Errorf("Expected ID to be '456', got '%s'", toy.ID)
	}
}

func TestCreateVideoGame(t *testing.T) {
	toy := createVideoGame("789", "Video Game", 29.99, "Description", "Category", 10, 18, "Brand", "0.5kg", "Vendor", []string{"Suitable"})

	// Add assertions here to validate the toy object

	// Example assertion:
	if toy.ID != "789" {
		t.Errorf("Expected ID to be '789', got '%s'", toy.ID)
	}
}
