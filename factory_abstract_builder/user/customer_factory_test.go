package user

import (
	"testing"
	"time"
)

func TestCustomerFactory_CreateUser(t *testing.T) {
	cf := NewCustomerFactory()

	id := "456"
	name := "Jane Smith"
	email := "jane.smith@example.com"
	contact := "9876543210"
	profilePicture := "profile.jpg"
	password := "password456"
	customerType := "regular"
	dob := time.Date(1995, time.February, 15, 0, 0, 0, 0, time.UTC)

	customer := cf.CreateUser(id, name, email, contact, profilePicture, password, customerType, dob)

	if customer.User.ID != id {
		t.Errorf("Expected ID to be %s, but got %s", id, customer.User.ID)
	}

	if customer.User.UserType != "customer" {
		t.Errorf("Expected UserType to be 'customer', but got %s", customer.User.UserType)
	}

	if customer.User.Name != name {
		t.Errorf("Expected Name to be %s, but got %s", name, customer.User.Name)
	}

	// Add more assertions for other fields if needed

	if customer.CustomerType != customerType {
		t.Errorf("Expected CustomerType to be %s, but got %s", customerType, customer.CustomerType)
	}
}
