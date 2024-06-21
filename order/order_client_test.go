package order

import (
	"testing"
)

func TestCreateOrder(t *testing.T) {
	items := []OrderItem{
		OrderItem{
			"id1",
			"item1",
			1,
			1343.12,
		},
		OrderItem{
			"id22",
			"item12",
			1,
			1343.12,
		},
		OrderItem{
			"id3",
			"item12",
			1,
			1343.12,
		},
		OrderItem{
			"id4",
			"item112",
			1,
			1343.12,
		},
	}

	_, err := createOrder("regular", "credit", "USD", "address", "name", "email", items)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

}
