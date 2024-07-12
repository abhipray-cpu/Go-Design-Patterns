package user

import (
	"testing"
	"time"
)

func TestAdminFactory_CreateUser(t *testing.T) {
	af := NewAdminFactory()

	id := "123"
	name := "John Doe"
	email := "john.doe@example.com"
	contact := "1234567890"
	profilePicture := "profile.jpg"
	password := "password123"
	permissions := []string{"create", "read", "update", "delete"}
	dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

	admin := af.CreateUser(id, name, email, contact, profilePicture, password, permissions, dob)

	if admin.User.ID != id {
		t.Errorf("Expected ID to be %s, but got %s", id, admin.User.ID)
	}

	if admin.User.UserType != "admin" {
		t.Errorf("Expected UserType to be 'admin', but got %s", admin.User.UserType)
	}

	if admin.User.Name != name {
		t.Errorf("Expected Name to be %s, but got %s", name, admin.User.Name)
	}

	// Add more assertions for other fields if needed

	if len(admin.Permissions) != len(permissions) {
		t.Errorf("Expected Permissions length to be %d, but got %d", len(permissions), len(admin.Permissions))
	}

	for i, p := range permissions {
		if admin.Permissions[i] != p {
			t.Errorf("Expected Permission at index %d to be %s, but got %s", i, p, admin.Permissions[i])
		}
	}
}
