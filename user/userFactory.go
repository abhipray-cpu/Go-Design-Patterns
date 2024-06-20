package user

import (
	"errors"
)

type UserType string

const (
	customer UserType = "customer"
	vendor   UserType = "vendor"
	admin    UserType = "admin"
)

// GetUser returns an instance of IUser based on the provided user type.
// It takes a UserType parameter and returns an IUser and an error.
// The error will be nil if the user type is valid, otherwise it will be an "invalid user type" error.
func GetUser(userType UserType) (IUser, error) {
	switch userType {
	case customer:
		return NewCustomerFactory(),nil
	case vendor:
		return NewVendorFactory(),nil
	case admin:
		return NewAdminFactory(), nil
	default:
		return nil, errors.New("invalid user type")
	}
}
