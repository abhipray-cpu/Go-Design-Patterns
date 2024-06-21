package health

import (
	"testing"
)

func TestCreateHairCareProduct(t *testing.T) {
	// Test case 1
	product := createHairCareProduct("1", "Shampoo", 9.99, "Cleans and nourishes hair", "ABC Company", "Hair Care", []string{"Water", "Sodium Laureth Sulfate"}, "BrandX", []string{"Apply to wet hair", "Massage into scalp"}, "Shampoo", "", "250ml", true, "Unisex")

	// Assert product properties
	if product.Id != "1" {
		t.Errorf("Expected product ID to be '1', got '%s'", product.Id)
	}
	if product.Name != "Shampoo" {
		t.Errorf("Expected product name to be 'Shampoo', got '%s'", product.Name)
	}
	if product.Price != 9.99 {
		t.Errorf("Expected product price to be 9.99, got %.2f", product.Price)
	}
	// Add more assertions here for other product properties

	// Test case 2
	// Add more test cases here
}

func TestCreateMakeupProduct(t *testing.T) {
	// Test case 1
	product := createMakeupProduct("1", "Lipstick", 14.99, "Long-lasting lipstick", "XYZ Company", "Makeup", []string{"Ricinus Communis Seed Oil", "Cera Alba"}, "BrandY", "Lipstick", "5g", "Red", "Matte")

	// Assert product properties
	if product.Id != "1" {
		t.Errorf("Expected product ID to be '1', got '%s'", product.Id)
	}
	if product.Name != "Lipstick" {
		t.Errorf("Expected product name to be 'Lipstick', got '%s'", product.Name)
	}
	if product.Price != 14.99 {
		t.Errorf("Expected product price to be 14.99, got %.2f", product.Price)
	}
	// Add more assertions here for other product properties

	// Test case 2
	// Add more test cases here
}

func TestCreatePersonalCareProduct(t *testing.T) {
	// Test case 1
	product := createPersonalCareProduct("1", "Toothpaste", 4.99, "Freshens breath and prevents cavities", "PQR Company", "Personal Care", []string{"Sorbitol", "Sodium Fluoride"}, "BrandZ", []string{"Brush teeth twice a day"}, "Toothpaste", "100g", "Unisex")

	// Assert product properties
	if product.Id != "1" {
		t.Errorf("Expected product ID to be '1', got '%s'", product.Id)
	}
	if product.Name != "Toothpaste" {
		t.Errorf("Expected product name to be 'Toothpaste', got '%s'", product.Name)
	}
	if product.Price != 4.99 {
		t.Errorf("Expected product price to be 4.99, got %.2f", product.Price)
	}
	// Add more assertions here for other product properties

	// Test case 2
	// Add more test cases here
}

func TestCreateSkinCareProduct(t *testing.T) {
	// Test case 1
	product := createSkinCareProduct("1", "Moisturizer", 19.99, "Hydrates and nourishes skin", "LMN Company", "Skin Care", []string{"Water", "Glycerin"}, "BrandA", []string{"Apply to clean face"}, "Moisturizer", "SPF 15", "50ml", "Unisex")

	// Assert product properties
	if product.Id != "1" {
		t.Errorf("Expected product ID to be '1', got '%s'", product.Id)
	}
	if product.Name != "Moisturizer" {
		t.Errorf("Expected product name to be 'Moisturizer', got '%s'", product.Name)
	}
	if product.Price != 19.99 {
		t.Errorf("Expected product price to be 19.99, got %.2f", product.Price)
	}
	// Add more assertions here for other product properties

	// Test case 2
	// Add more test cases here
}
