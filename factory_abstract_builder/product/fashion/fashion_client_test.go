package fashion

import (
	"testing"
)

func TestCreateBagProduct(t *testing.T) {
	// Test case 1
	product := createBagProduct("1", "Bag 1", 99.99, "Description 1", "Vendor 1", "Brand 1", "Category 1", "Material 1", "Size 1", "Color 1", "Female", "Style 1", "Product Type 1")

	// Assertion 1
	if product.Id != "1" {
		t.Errorf("Expected Id to be '1', but got '%s'", product.Id)
	}

	// Assertion 2
	if product.Name != "Bag 1" {
		t.Errorf("Expected Name to be 'Bag 1', but got '%s'", product.Name)
	}

	// Assertion 3
	if product.Price != 99.99 {
		t.Errorf("Expected Price to be 99.99, but got %.2f", product.Price)
	}

	// Assertion 4
	if product.Description != "Description 1" {
		t.Errorf("Expected Description to be 'Description 1', but got '%s'", product.Description)
	}

	// Assertion 5
	if product.Vendor != "Vendor 1" {
		t.Errorf("Expected Vendor to be 'Vendor 1', but got '%s'", product.Vendor)
	}

	// Assertion 6
	if product.Brand != "Brand 1" {
		t.Errorf("Expected Brand to be 'Brand 1', but got '%s'", product.Brand)
	}

	// Assertion 7
	if product.Category != "Category 1" {
		t.Errorf("Expected Category to be 'Category 1', but got '%s'", product.Category)
	}

	// Assertion 8
	if product.Material != "Material 1" {
		t.Errorf("Expected Material to be 'Material 1', but got '%s'", product.Material)
	}

	// Assertion 9
	if product.Size != "Size 1" {
		t.Errorf("Expected Size to be 'Size 1', but got '%s'", product.Size)
	}

	// Assertion 10
	if product.Color != "Color 1" {
		t.Errorf("Expected Color to be 'Color 1', but got '%s'", product.Color)
	}

	// Assertion 12
	if product.Style != "Style 1" {
		t.Errorf("Expected Style to be 'Style 1', but got '%s'", product.Style)
	}
}

func TestCreateJewelryProduct(t *testing.T) {
	// Test case 1
	product := createJewelryProduct("1", "Jewelry 1", 199.99, "Description 1", "Vendor 1", "Brand 1", "Category 1", "Material 1", "Size 1", "Color 1", "Gender 1", "Style 1", "Product Type 1", "Gemstone 1")

	// Assertion 1
	if product.Id != "1" {
		t.Errorf("Expected Id to be '1', but got '%s'", product.Id)
	}

	// Assertion 2
	if product.Name != "Jewelry 1" {
		t.Errorf("Expected Name to be 'Jewelry 1', but got '%s'", product.Name)
	}

	// Assertion 3
	if product.Price != 199.99 {
		t.Errorf("Expected Price to be 199.99, but got %.2f", product.Price)
	}

	// Assertion 4
	if product.Description != "Description 1" {
		t.Errorf("Expected Description to be 'Description 1', but got '%s'", product.Description)
	}

	// Assertion 5
	if product.Vendor != "Vendor 1" {
		t.Errorf("Expected Vendor to be 'Vendor 1', but got '%s'", product.Vendor)
	}

	// Assertion 6
	if product.Brand != "Brand 1" {
		t.Errorf("Expected Brand to be 'Brand 1', but got '%s'", product.Brand)
	}

	// Assertion 7
	if product.Category != "Category 1" {
		t.Errorf("Expected Category to be 'Category 1', but got '%s'", product.Category)
	}

	// Assertion 8
	if product.Material != "Material 1" {
		t.Errorf("Expected Material to be 'Material 1', but got '%s'", product.Material)
	}

	// Assertion 9
	if product.Size != "Size 1" {
		t.Errorf("Expected Size to be 'Size 1', but got '%s'", product.Size)
	}

	// Assertion 10
	if product.Color != "Color 1" {
		t.Errorf("Expected Color to be 'Color 1', but got '%s'", product.Color)
	}

	// Assertion 11
	if product.Gender != "Gender 1" {
		t.Errorf("Expected Gender to be 'Gender 1', but got '%s'", product.Gender)
	}

	// Assertion 12
	if product.Style != "Style 1" {
		t.Errorf("Expected Style to be 'Style 1', but got '%s'", product.Style)
	}

	// Assertion 14
	if product.GemStone != "Gemstone 1" {
		t.Errorf("Expected GemStone to be 'Gemstone 1', but got '%s'", product.GemStone)
	}

	// Test case 2
	// Add more test cases here...
}

func TestCreateShoeProduct(t *testing.T) {
	// Test case 1
	product := createShoeProduct("1", "Shoe 1", 149.99, "Description 1", "Vendor 1", "Brand 1", "Category 1", "Material 1", "Size 1", "Color 1", "Gender 1", "Style 1", "Product Type 1")

	// Assertion 1
	if product.Id != "1" {
		t.Errorf("Expected Id to be '1', but got '%s'", product.Id)
	}

	// Assertion 2
	if product.Name != "Shoe 1" {
		t.Errorf("Expected Name to be 'Shoe 1', but got '%s'", product.Name)
	}

	// Assertion 3
	if product.Price != 149.99 {
		t.Errorf("Expected Price to be 149.99, but got %.2f", product.Price)
	}

	// Assertion 4
	if product.Description != "Description 1" {
		t.Errorf("Expected Description to be 'Description 1', but got '%s'", product.Description)
	}

	// Assertion 5
	if product.Vendor != "Vendor 1" {
		t.Errorf("Expected Vendor to be 'Vendor 1', but got '%s'", product.Vendor)
	}

	// Assertion 6
	if product.Brand != "Brand 1" {
		t.Errorf("Expected Brand to be 'Brand 1', but got '%s'", product.Brand)
	}

	// Assertion 7
	if product.Category != "Category 1" {
		t.Errorf("Expected Category to be 'Category 1', but got '%s'", product.Category)
	}

	// Assertion 8
	if product.Material != "Material 1" {
		t.Errorf("Expected Material to be 'Material 1', but got '%s'", product.Material)
	}

	// Assertion 9
	if product.Size != "Size 1" {
		t.Errorf("Expected Size to be 'Size 1', but got '%s'", product.Size)
	}

	// Assertion 10
	if product.Color != "Color 1" {
		t.Errorf("Expected Color to be 'Color 1', but got '%s'", product.Color)
	}

	// Assertion 11
	if product.Gender != "Gender 1" {
		t.Errorf("Expected Gender to be 'Gender 1', but got '%s'", product.Gender)
	}

	// Assertion 12
	if product.Style != "Style 1" {
		t.Errorf("Expected Style to be 'Style 1', but got '%s'", product.Style)
	}

	// Test case 2
	// Add more test cases here...
}

func TestCreateClothingProduct(t *testing.T) {
	// Test case 1
	product := createClothingProduct("1", "Clothing 1", 79.99, "Description 1", "Vendor 1", "Brand 1", "Category 1", "Material 1", "Size 1", "Color 1", "Gender 1", "Style 1", "Product Type 1")

	// Assertion 1
	if product.Id != "1" {
		t.Errorf("Expected Id to be '1', but got '%s'", product.Id)
	}

	// Assertion 2
	if product.Name != "Clothing 1" {
		t.Errorf("Expected Name to be 'Clothing 1', but got '%s'", product.Name)
	}

	// Assertion 3
	if product.Price != 79.99 {
		t.Errorf("Expected Price to be 79.99, but got %.2f", product.Price)
	}

	// Assertion 4
	if product.Description != "Description 1" {
		t.Errorf("Expected Description to be 'Description 1', but got '%s'", product.Description)
	}

	// Assertion 5
	if product.Vendor != "Vendor 1" {
		t.Errorf("Expected Vendor to be 'Vendor 1', but got '%s'", product.Vendor)
	}

	// Assertion 6
	if product.Brand != "Brand 1" {
		t.Errorf("Expected Brand to be 'Brand 1', but got '%s'", product.Brand)
	}

	// Assertion 7
	if product.Category != "Category 1" {
		t.Errorf("Expected Category to be 'Category 1', but got '%s'", product.Category)
	}

	// Assertion 8
	if product.Material != "Material 1" {
		t.Errorf("Expected Material to be 'Material 1', but got '%s'", product.Material)
	}

	// Assertion 9
	if product.Size != "Size 1" {
		t.Errorf("Expected Size to be 'Size 1', but got '%s'", product.Size)
	}

	// Assertion 10
	if product.Color != "Color 1" {
		t.Errorf("Expected Color to be 'Color 1', but got '%s'", product.Color)
	}

	// Assertion 11
	if product.Gender != "Gender 1" {
		t.Errorf("Expected Gender to be 'Gender 1', but got '%s'", product.Gender)
	}

	// Assertion 12
	if product.Style != "Style 1" {
		t.Errorf("Expected Style to be 'Style 1', but got '%s'", product.Style)
	}

	// Test case 2
	// Add more test cases here...
}

func TestCreateWatchProduct(t *testing.T) {
	// Test case 1
	product := createWatchProduct("1", "Watch 1", 299.99, "Description 1", "Vendor 1", "Brand 1", "Category 1", "Material 1", "Size 1", "Color 1", "Gender 1", "Style 1", "Product Type 1", "Resistance 1")

	// Assertion 1
	if product.Id != "1" {
		t.Errorf("Expected Id to be '1', but got '%s'", product.Id)
	}

	// Assertion 2
	if product.Name != "Watch 1" {
		t.Errorf("Expected Name to be 'Watch 1', but got '%s'", product.Name)
	}

	// Assertion 3
	if product.Price != 299.99 {
		t.Errorf("Expected Price to be 299.99, but got %.2f", product.Price)
	}

	// Assertion 4
	if product.Description != "Description 1" {
		t.Errorf("Expected Description to be 'Description 1', but got '%s'", product.Description)
	}

	// Assertion 5
	if product.Vendor != "Vendor 1" {
		t.Errorf("Expected Vendor to be 'Vendor 1', but got '%s'", product.Vendor)
	}

	// Assertion 6
	if product.Brand != "Brand 1" {
		t.Errorf("Expected Brand to be 'Brand 1', but got '%s'", product.Brand)
	}

	// Assertion 7
	if product.Category != "Category 1" {
		t.Errorf("Expected Category to be 'Category 1', but got '%s'", product.Category)
	}

	// Assertion 8
	if product.Material != "Material 1" {
		t.Errorf("Expected Material to be 'Material 1', but got '%s'", product.Material)
	}

	// Assertion 9
	if product.Size != "Size 1" {
		t.Errorf("Expected Size to be 'Size 1', but got '%s'", product.Size)
	}

	// Assertion 10
	if product.Color != "Color 1" {
		t.Errorf("Expected Color to be 'Color 1', but got '%s'", product.Color)
	}

	// Assertion 11
	if product.Gender != "Gender 1" {
		t.Errorf("Expected Gender to be 'Gender 1', but got '%s'", product.Gender)
	}

	// Assertion 12
	if product.Style != "Style 1" {
		t.Errorf("Expected Style to be 'Style 1', but got '%s'", product.Style)
	}

	// Test case 2
	// Add more test cases here...
}

// Test case 2
// Add more test cases here...
