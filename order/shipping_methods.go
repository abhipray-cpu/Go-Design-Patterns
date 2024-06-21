package order

import (
	"errors"
	"time"
)

// ShippingMethod interface defines the behavior of shipping methods.
type ShippingMethod interface {
	GetCost() float64
	GetExpectedDelivery() time.Time
}

// ShippingDetails contains details about the shipping method and address.
type ShippingDetails struct {
	Method           string    // Shipping method, e.g., "Standard", "Express"
	Address          string    // Shipping address
	Cost             float64   // Cost of shipping - corrected to float64 for consistency and calculations
	ExpectedDelivery time.Time // Estimated delivery date
}

// Ship sets or updates shipping details and returns the updated object.
// In a real application, this might involve more complex logic such as API calls.
func (s *ShippingDetails) Ship(method, address string, items []OrderItem) (ShippingDetails, error) {
	if address == "" {
		return ShippingDetails{}, errors.New("Address is required")
	}
	if len(items) == 0 {
		return ShippingDetails{}, errors.New("Items are required")

	}

	cost := s.GetCost(len(items), method)
	delivery := s.GetExpectedDelivery(method)
	return ShippingDetails{
		Method:           method,
		Address:          address,
		Cost:             cost,
		ExpectedDelivery: delivery,
	}, nil
}

// GetCost returns the cost of shipping.
func (s *ShippingDetails) GetCost(items int, method string) float64 {
	var costPerItem float64
	var baseCost float64

	switch method {
	case "Standard":
		costPerItem = 2.0 // Example cost per item for standard shipping
		baseCost = 5.0    // Example base cost for standard shipping
	case "Express":
		costPerItem = 3.5 // Example cost per item for express shipping
		baseCost = 10.0   // Example base cost for express shipping
	case "Overnight":
		costPerItem = 5.0 // Example cost per item for overnight shipping
		baseCost = 15.0   // Example base cost for overnight shipping
	default:
		costPerItem = 2.0 // Default cost per item
		baseCost = 5.0    // Default base cost
	}

	totalCost := baseCost + (float64(items) * costPerItem)
	return totalCost
}

func (s *ShippingDetails) GetExpectedDelivery(method string) time.Time {
	var daysToAdd int

	switch method {
	case "Standard":
		daysToAdd = 5 // Example delivery time for standard shipping
	case "Express":
		daysToAdd = 2 // Example delivery time for express shipping
	case "Overnight":
		daysToAdd = 1 // Example delivery time for overnight shipping
	default:
		daysToAdd = 5 // Default delivery time
	}

	// Assuming s.OrderDate is the date when the order was placed
	expectedDelivery := time.Now().AddDate(0, 0, daysToAdd)
	return expectedDelivery
}
