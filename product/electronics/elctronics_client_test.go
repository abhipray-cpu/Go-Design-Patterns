package electronics

import (
	"testing"
)

// Example assertion:
func TestCreateMobileProduct(t *testing.T) {
	id := "123"
	name := "iPhone 12"
	price := 999.99
	description := "The latest iPhone model"
	vendor := "Apple"
	brand := "Apple"
	category := "Mobile"
	screenSize := "6.1 inches"
	batteryLife := "Up to 17 hours"
	os := "iOS"
	storage := "256GB"
	cpu := "A14 Bionic"
	ram := "4GB"
	resolution := "2532 x 1170 pixels"
	prodType := "Smartphone"
	weight := "164 grams"
	dimension := "146.7 x 71.5 x 7.4 mm"
	features := []string{"Face ID", "Dual-camera system", "Water and dust resistant"}

	product := createMobileProduct(id, name, price, description, vendor, brand, category, screenSize, batteryLife, os, storage, cpu, ram, resolution, prodType, weight, dimension, features)

	// Add assertions here to validate the created mobile product

	// Example assertions:
	if product.ID != id {
		t.Errorf("Expected ID to be %s, but got %s", id, product.ID)
	}

	if product.Name != name {
		t.Errorf("Expected Name to be %s, but got %s", name, product.Name)
	}

	if product.Price != price {
		t.Errorf("Expected Price to be %f, but got %f", price, product.Price)
	}

	if product.Description != description {
		t.Errorf("Expected Description to be %s, but got %s", description, product.Description)
	}

	if product.Vendor != vendor {
		t.Errorf("Expected Vendor to be %s, but got %s", vendor, product.Vendor)
	}

	if product.Brand != brand {
		t.Errorf("Expected Brand to be %s, but got %s", brand, product.Brand)
	}

	if product.Category != category {
		t.Errorf("Expected Category to be %s, but got %s", category, product.Category)
	}

	if product.ScreenSize != screenSize {
		t.Errorf("Expected ScreenSize to be %s, but got %s", screenSize, product.ScreenSize)
	}

	if product.BatteryLife != batteryLife {
		t.Errorf("Expected BatteryLife to be %s, but got %s", batteryLife, product.BatteryLife)
	}

	if product.OS != os {
		t.Errorf("Expected OS to be %s, but got %s", os, product.OS)
	}

	if product.StorageCapacity != storage {
		t.Errorf("Expected Storage to be %s, but got %s", storage, product.StorageCapacity)
	}

	if product.CPU != cpu {
		t.Errorf("Expected CPU to be %s, but got %s", cpu, product.CPU)
	}

	if product.RAM != ram {
		t.Errorf("Expected RAM to be %s, but got %s", ram, product.RAM)
	}

	if product.Resolution != resolution {
		t.Errorf("Expected Resolution to be %s, but got %s", resolution, product.Resolution)
	}

	if product.Type != prodType {
		t.Errorf("Expected ProdType to be %s, but got %s", prodType, product.Type)
	}

	if product.Weight != weight {
		t.Errorf("Expected Weight to be %s, but got %s", weight, product.Weight)
	}

	if product.Dimensions != dimension {
		t.Errorf("Expected Dimension to be %s, but got %s", dimension, product.Dimensions)
	}

	if len(product.Features) != len(features) {
		t.Errorf("Expected %d features, but got %d", len(features), len(product.Features))
	} else {
		for i, feature := range features {
			if product.Features[i] != feature {
				t.Errorf("Expected feature at index %d to be %s, but got %s", i, feature, product.Features[i])
			}
		}
	}
}
