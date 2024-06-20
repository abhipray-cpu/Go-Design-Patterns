package user

import "time"

type AdminFactory struct {
	IUser
}

// Admin represents an administrative user with specific permissions.
type Admin struct {
	User        User     // The underlying user information for the admin.
	Permissions []string // The list of permissions assigned to the admin.
}

// NewAdminFactory creates a new instance of AdminFactory.
func NewAdminFactory() *AdminFactory {
	return &AdminFactory{}
}

// CreateUser creates a new Admin user with the specified details and returns a pointer to the created Admin object.
func (af *AdminFactory) CreateUser(id, name, email, contact, profilePicture, password string, permissions []string, dob time.Time) *Admin {
	return &Admin{
		User: User{
			ID:             id,
			UserType:       "admin",
			Name:           name,
			Email:          email,
			Contact:        contact,
			ProfilePicture: profilePicture,
			Password:       password,
			Status:         "active",
			LastLogin:      time.Now(),
			DOB:            dob,
		},
		Permissions: permissions,
	}
}
