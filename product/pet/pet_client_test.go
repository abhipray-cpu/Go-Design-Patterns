package pet

import (
	"testing"
)

func TestCreateAccessory(t *testing.T) {
	// Test case 1
	accessory := createAccessory("1", "Collar", 9.99, "A stylish collar for pets", "Vendor A", "Accessories", "Brand X", []string{"Dog", "Cat"}, "Medium", "0.2 lbs", "Nylon", "Medium", "Accessory")
	if accessory.ID != "1" {
		t.Errorf("Expected ID: 1, Got: %s", accessory.ID)
	}
	// Add more test cases here...
}

func TestCreateFood(t *testing.T) {
	// Test case 1
	food := createFood("2", "Dog Food", 19.99, "Healthy food for dogs", "Vendor B", "Food", "Brand Y", []string{"Chicken", "Rice"}, "High in protein", "5 lbs", "Large", "Food", []string{"Dog"})
	if food.Price != 19.99 {
		t.Errorf("Expected Price: 19.99, Got: %.2f", food.Price)
	}
	// Add more test cases here...
}

func TestCreateGrooming(t *testing.T) {
	// Test case 1
	grooming := createGrooming("3", "Shampoo", 7.99, "Gentle shampoo for pets", "Vendor C", "Grooming", "Brand Z", []string{"Dog", "Cat"}, "Long-lasting", "500ml", "Medium", "Grooming", []string{"Water", "Coconut oil"})
	if grooming.Description != "Gentle shampoo for pets" {
		t.Errorf("Expected Description: Gentle shampoo for pets, Got: %s", grooming.Description)
	}
	// Add more test cases here...
}

func TestCreateToy(t *testing.T) {
	// Test case 1
	toy := createToy("4", "Ball", 5.99, "A fun toy for pets", "Vendor D", "Toys", "Brand W", []string{"Dog", "Cat"}, "Durable", "Rubber", "Toy", "Small")
	if toy.Vendor != "Vendor D" {
		t.Errorf("Expected Vendor: Vendor D, Got: %s", toy.Vendor)
	}
	// Add more test cases here...
}
