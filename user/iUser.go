package user

import "time"

// IUser represents the interface for a user in the system.
type IUser interface {
	// GetID returns the unique identifier of the user.
	GetID() string

	// GetUserType returns the type of the user.
	GetUserType() string

	// GetName returns the name of the user.
	GetName() string

	// GetEmail returns the email address of the user.
	GetEmail() string

	// GetContact returns the contact number of the user.
	GetContact() string

	// GetProfilePicture returns the URL of the user's profile picture.
	GetProfilePicture() string

	// GetPassword returns the password of the user.
	GetPassword() string

	// GetStatus returns the status of the user.
	GetStatus() string

	// GetLastLogin returns the timestamp of the user's last login.
	GetLastLogin() time.Time

	// GetDOB returns the date of birth of the user.
	GetDOB() time.Time
}