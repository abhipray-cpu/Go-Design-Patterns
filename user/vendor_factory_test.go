package user

import (
	"testing"
	"time"
)

func TestVendorFactory_CreateUser(t *testing.T) {
	vf := NewVendorFactory()

	id := "123"
	name := "John Doe"
	email := "john.doe@example.com"
	contact := "1234567890"
	profilePicture := "profile.jpg"
	brand := "ACME"
	business := "ACME Corporation"
	password := "password123"
	dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

	vendor := vf.CreateUser(id, name, email, contact, profilePicture, brand, business, password, dob)

	if vendor.User.ID != id {
		t.Errorf("Expected ID to be %s, but got %s", id, vendor.User.ID)
	}

	if vendor.User.UserType != "vendor" {
		t.Errorf("Expected UserType to be 'vendor', but got %s", vendor.User.UserType)
	}

	if vendor.User.Name != name {
		t.Errorf("Expected Name to be %s, but got %s", name, vendor.User.Name)
	}

	// Add more assertions for other fields if needed

	if vendor.Brand != brand {
		t.Errorf("Expected Brand to be %s, but got %s", brand, vendor.Brand)
	}

	if vendor.Business != business {
		t.Errorf("Expected Business to be %s, but got %s", business, vendor.Business)
	}
}
