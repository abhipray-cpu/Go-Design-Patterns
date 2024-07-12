package user

import "time"

type VendorFactory struct {
	IUser
}

// Vendor represents a vendor in the e-commerce system.
type Vendor struct {
	User     User   // The user associated with the vendor.
	Brand    string // The brand name of the vendor.
	Business string // The business name of the vendor.
}

// NewVendorFactory creates a new instance of VendorFactory.
func NewVendorFactory() *VendorFactory {
	return &VendorFactory{}
}

// CreateUser creates a new Vendor user with the given information and returns a pointer to the Vendor object.
func (vf *VendorFactory) CreateUser(id, name, email, contact, profilePicture, brand string, business string, password string, dob time.Time) *Vendor {
	return &Vendor{
		User: User{
			ID:             id,
			UserType:       "vendor",
			Name:           name,
			Email:          email,
			Contact:        contact,
			ProfilePicture: profilePicture,
			Password:       password,
			Status:         "active",
			LastLogin:      time.Now(),
			DOB:            dob,
		},
		Brand:    brand,
		Business: business,
	}
}
