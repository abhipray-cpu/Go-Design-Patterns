package home

import (
	"testing"
)

// Add assertions here to verify the created product
func TestCreateBathProduct(t *testing.T) {
	product := createBathProduct("123", "Bath Product", 9.99, "Description", "Vendor", "Category", "Brand", "Product Type", "Volume", "Scent", "Skin Type", "Material", "Dimension", "Weight", "Color", "Capacity", "Style", true)

	// Add assertions here to verify the created product
	if product.ID != "123" {
		t.Errorf("Expected ID to be '123', but got '%s'", product.ID)
	}
	if product.Name != "Bath Product" {
		t.Errorf("Expected Name to be 'Bath Product', but got '%s'", product.Name)
	}
	if product.Price != 9.99 {
		t.Errorf("Expected Price to be 9.99, but got %.2f", product.Price)
	}
	// Add more assertions for other properties
}

func TestCreateDecorProduct(t *testing.T) {
	product := createDecorProduct("456", "Decor Product", 19.99, "Description", "Vendor", "Category", "Brand", "Product Type", "Color", "Material", "Weight", "Style", "Dimension", true)

	// Add assertions here to verify the created product
	if product.ID != "456" {
		t.Errorf("Expected ID to be '456', but got '%s'", product.ID)
	}
	if product.Name != "Decor Product" {
		t.Errorf("Expected Name to be 'Decor Product', but got '%s'", product.Name)
	}
	if product.Price != 19.99 {
		t.Errorf("Expected Price to be 19.99, but got %.2f", product.Price)
	}
	// Add more assertions for other properties
}

func TestCreateFurnitureProduct(t *testing.T) {
	product := createFurnitureProduct("789", "Furniture Product", 29.99, "Description", "Vendor", "Category", "Brand", "Product Type", "Color", "Material", "Weight", "Style", "Dimension", true, "Capacity")

	// Add assertions here to verify the created product
	if product.ID != "789" {
		t.Errorf("Expected ID to be '789', but got '%s'", product.ID)
	}
	if product.Name != "Furniture Product" {
		t.Errorf("Expected Name to be 'Furniture Product', but got '%s'", product.Name)
	}
	if product.Price != 29.99 {
		t.Errorf("Expected Price to be 29.99, but got %.2f", product.Price)
	}
	// Add more assertions for other properties
}

func TestCreateKitchenProduct(t *testing.T) {
	product := createKitchenProduct("101112", "Kitchen Product", 39.99, "Description", "Vendor", "Category", "Brand", "Product Type", "Color", "Material", "Weight", "Style", "Dimension", "Capacity", true, true, "Power")

	// Add assertions here to verify the created product
	if product.ID != "101112" {
		t.Errorf("Expected ID to be '101112', but got '%s'", product.ID)
	}
	if product.Name != "Kitchen Product" {
		t.Errorf("Expected Name to be 'Kitchen Product', but got '%s'", product.Name)
	}
	if product.Price != 39.99 {
		t.Errorf("Expected Price to be 39.99, but got %.2f", product.Price)
	}
	// Add more assertions for other properties
}

func TestCreateToolProduct(t *testing.T) {
	product := createToolProduct("131415", "Tool Product", 49.99, "Description", "Vendor", "Category", "Brand", "Product Type", "Color", "Material", "Weight", "Style", "Dimensions", "Power Source")

	// Add assertions here to verify the created product
	if product.ID != "131415" {
		t.Errorf("Expected ID to be '131415', but got '%s'", product.ID)
	}
	if product.Name != "Tool Product" {
		t.Errorf("Expected Name to be 'Tool Product', but got '%s'", product.Name)
	}
	if product.Price != 49.99 {
		t.Errorf("Expected Price to be 49.99, but got %.2f", product.Price)
	}
	// Add more assertions for other properties
}
