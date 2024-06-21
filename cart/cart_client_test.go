package cart

import (
	"testing"
)

func TestCreateCart(t *testing.T) {
	cartBuilder := createCart("regular", "1", "item1", "item1", 100.0, 1)
	cart := cartBuilder.Build()
	cartBuilder.SetShippingDetails("123", "123", "123", "123", "123")
	cartBuilder.SetPaymentDetails("123", "123", "123")
	if cart.User != "1" {
		t.Errorf("Expected User to be 'regular', but got '%s'", cart.User)
	}

	if _, ok := cart.Items["1"]; ok {
		item := cart.Items["1"]
		if item.Name != "item1" {
			t.Errorf("Expected item Name to be 'item1', but got '%s'", item.Name)
		}
		if item.Description != "item1" {
			t.Errorf("Expected item Description to be 'item1', but got '%s'", item.Description)
		}
		if item.Price != 100.0 {
			t.Errorf("Expected item Price to be 100.0, but got %.2f", item.Price)
		}
		if item.Quantity != 1 {
			t.Errorf("Expected item Quantity to be 1, but got %d", item.Quantity)
		}
	} else {
		t.Errorf("Item with id '1' not found in Items map")
	}

}

func TestCreatePremiumCart(t *testing.T) {
	cartBuilder := createCart("premium", "2", "item2", "item2", 200.0, 2)
	cart := cartBuilder.Build()
	cartBuilder.SetShippingDetails("456", "456", "456", "456", "456")
	cartBuilder.SetPaymentDetails("456", "456", "456")
	if cart.User != "2" {
		t.Errorf("Expected User to be 'premium', but got '%s'", cart.User)
	}

	if _, ok := cart.Items["2"]; ok {
		item := cart.Items["2"]
		if item.Name != "item2" {
			t.Errorf("Expected item Name to be 'item2', but got '%s'", item.Name)
		}
		if item.Description != "item2" {
			t.Errorf("Expected item Description to be 'item2', but got '%s'", item.Description)
		}
		if item.Price != 200.0 {
			t.Errorf("Expected item Price to be 200.0, but got %.2f", item.Price)
		}
		if item.Quantity != 2 {
			t.Errorf("Expected item Quantity to be 2, but got %d", item.Quantity)
		}
	} else {
		t.Errorf("Item with id '2' not found in Items map")
	}
}
