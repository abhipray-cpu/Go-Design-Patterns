package sports

import (
	"testing"
)

func TestCreateCampingProduct(t *testing.T) {
	// Test case 1
	product := createCampingProduct("123", "Camping Tent", 99.99, "A spacious tent for camping", "Camping", 8, 99, "10x10 ft", "High", "Nylon", "5 lbs", "ABC Company", "Tent", "Large", "Outdoor", "Green", []string{"Family", "Camping"})
	if product.ID != "123" {
		t.Errorf("Expected product ID to be '123', but got '%s'", product.ID)
	}
	// Add more test cases here...
}

func TestCreateEquipmentProduct(t *testing.T) {
	// Test case 1
	product := createEquipmentProduct("456", "Basketball", 29.99, "A basketball for outdoor use", "Equipment", 6, 99, "9.5 in", "Medium", "Rubber", "1.5 lbs", "XYZ Company", "Ball", "Standard", "Nike", "Orange", []string{"Outdoor", "Basketball"})
	if product.Price != 29.99 {
		t.Errorf("Expected product price to be 29.99, but got %.2f", product.Price)
	}
	// Add more test cases here...
}

func TestCreateExerciseProduct(t *testing.T) {
	// Test case 1
	product := createExerciseProduct("789", "Yoga Mat", 19.99, "A comfortable mat for yoga exercises", "Exercise", 12, 99, "24x68 in", "High", "PVC", "2 lbs", "PQR Company", "Mat", "Standard", "Yoga", "Purple", []string{"Yoga", "Exercise"})
	if product.Description != "A comfortable mat for yoga exercises" {
		t.Errorf("Expected product description to be 'A comfortable mat for yoga exercises', but got '%s'", product.Description)
	}
	// Add more test cases here...
}

func TestCreateRecreationProduct(t *testing.T) {
	// Test case 1
	product := createRecreationProduct("987", "Frisbee", 9.99, "A flying disc for outdoor recreation", "Recreation", 6, 99, "9 in", "Medium", "Plastic", "0.5 lbs", "LMN Company", "Disc", "Standard", "Frisbee", "Red", []string{"Outdoor", "Recreation"})
	if product.Category != "Recreation" {
		t.Errorf("Expected product category to be 'Recreation', but got '%s'", product.Category)
	}
	// Add more test cases here...
}
