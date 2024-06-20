package user

import "time"

// User represents a user in the system.
type User struct {
	ID             string        // ID is the unique identifier for the user.
	UserType       string        // UserType represents the type of user (e.g., admin, customer).
	Name           string        // Name is the full name of the user.
	Email          string        // Email is the email address of the user.
	Contact        string        // Contact is the contact number of the user.
	ProfilePicture string        // ProfilePicture is the URL of the user's profile picture.
	Password       string        // Password is the hashed password of the user.
	Status         string        // Status represents the current status of the user (e.g., active, inactive).
	LastLogin      time.Time     // LastLogin is the timestamp of the user's last login.
	DOB            time.Time     // DOB is the date of birth of the user.
}
// GetID returns the id of the user.
func (u *User) GetID() string             { return u.ID }
// GetUserType returns the type of the user.
func (u *User) GetUserType() string       { return u.UserType }
// GetName returns the name of the user.
func (u *User) GetName() string           { return u.Name }
// GetEmail returns the email of the user.
func (u *User) GetEmail() string          { return u.Email }
// GetContact returns the contact of the user.
func (u *User) GetContact() string        { return u.Contact }
// GetProfilePicture returns the profile picture of the user.
func (u *User) GetProfilePicture() string { return u.ProfilePicture }
// GetPassword returns the password of the user.
func (u *User) GetPassword() string       { return u.Password }
// GetStatus returns the status of the user.
func (u *User) GetStatus() string         { return u.Status }
// GetLastLogin returns the the last login time of the user.
func (u *User) GetLastLogin() time.Time   { return u.LastLogin }
// GetDOB returns the DOB of the user.
func (u *User) GetDOB() time.Time         { return u.DOB }
