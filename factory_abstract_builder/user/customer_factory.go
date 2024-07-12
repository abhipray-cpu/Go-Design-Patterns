package user

import "time"

type CustomerFactory struct {
	IUser
}

// Customer represents a user who is also a customer.
type Customer struct {
	User                // Embedded User struct
	CustomerType string // The type of customer (e.g.,not registered, regular, premium)
}

// NewCustomerFactory creates a new instance of CustomerFactory.
func NewCustomerFactory() *CustomerFactory {
	return &CustomerFactory{}
}

// CreateUser creates a new Customer object with the given parameters and returns a pointer to it.
// The created Customer object is initialized with the provided id, name, email, contact, profilePicture,
// password, customer_type, and dob values. The created Customer object also inherits the User fields
// from the User struct, including the id, role, name, email, contact, profilePicture, password, status,
// created_at, and dob fields.
func (cf *CustomerFactory) CreateUser(id, name, email, contact, profilePicture, password string, customer_type string, dob time.Time) *Customer {
	return &Customer{
		User{
			id,
			"customer",
			name,
			email,
			contact,
			profilePicture,
			password,
			"active",
			time.Now(),
			dob,
		},
		customer_type,
	}
}
