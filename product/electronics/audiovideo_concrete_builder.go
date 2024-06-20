// Package electronics provides functionality for creating audio and video electronic products.
package electronics

// AudioVideoBuilder is a builder for creating audio and video electronic products.
type AudioVideoBuilder struct {
	electronic Electronic
}

// newAudioVideoBuilder creates a new instance of AudioVideoBuilder.
func newAudioVideoBuilder() *AudioVideoBuilder {
	return &AudioVideoBuilder{}
}

// SetAttributes sets the attributes of the electronic product.
func (e *AudioVideoBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	e.electronic.ID = id
	e.electronic.Name = name
	e.electronic.Price = price
	e.electronic.Description = description
	e.electronic.Vendor = vendor
}

// SetBrand sets the brand of the electronic product.
func (e *AudioVideoBuilder) SetBrand(brand string) {
	e.electronic.Brand = brand
}

// SetCategory sets the category of the electronic product.
func (e *AudioVideoBuilder) SetCategory(category string) {
	e.electronic.Category = category
}

// SetBatteryLife sets the battery life of the electronic product.
func (e *AudioVideoBuilder) SetBatteryLife(batteryLife string) {
	e.electronic.BatteryLife = batteryLife
}

// SetType sets the type of the electronic product.
func (e *AudioVideoBuilder) SetType(deviceType string) {
	e.electronic.Type = deviceType
}

// SetConnectivity sets the connectivity of the electronic product.
func (e *AudioVideoBuilder) SetConnectivity(connectivity string) {
	e.electronic.Connectivity = connectivity
}

// SetPowerOutput sets the power output of the electronic product.
func (e *AudioVideoBuilder) SetPowerOutput(powerOutput string) {
	e.electronic.PowerOutput = powerOutput
}

// SetWeight sets the weight of the electronic product.
func (e *AudioVideoBuilder) SetWeight(weight string) {
	e.electronic.Weight = weight
}

// SetDimensions sets the dimensions of the electronic product.
func (e *AudioVideoBuilder) SetDimensions(dimensions string) {
	e.electronic.Dimensions = dimensions
}

// SetFeatures sets the features of the electronic product.
func (e *AudioVideoBuilder) SetFeatures(features []string) {
	e.electronic.Features = features
}

// GetElectronic returns the created electronic product.
func (e *AudioVideoBuilder) GetElectronic() Electronic {
	return e.electronic
}

// The following functions are unused and exposed.

// SetScreenSize sets the screen size of the electronic product.
func (e *AudioVideoBuilder) SetScreenSize(screenSize string) {}

// SetResolution sets the resolution of the electronic product.
func (e *AudioVideoBuilder) SetResolution(resolution string) {}

// SetSensorType sets the sensor type of the electronic product.
func (e *AudioVideoBuilder) SetSensorType(sensorType string) {}

// SetRefreshRate sets the refresh rate of the electronic product.
func (e *AudioVideoBuilder) SetRefreshRate(refreshRate string) {}

// SetOpticalZoom sets the optical zoom of the electronic product.
func (e *AudioVideoBuilder) SetOpticalZoom(opticalZoom string) {}

// SetVideoQuality sets the video quality of the electronic product.
func (e *AudioVideoBuilder) SetVideoQuality(videoQuality string) {}

// SetEnergyRating sets the energy rating of the electronic product.
func (e *AudioVideoBuilder) SetEnergyRating(energyRating string) {}

// SetCapacity sets the capacity of the electronic product.
func (e *AudioVideoBuilder) SetCapacity(capacity string) {}

// SetCPU sets the CPU of the electronic product.
func (e *AudioVideoBuilder) SetCPU(cpu string) {}

// SetRAM sets the RAM of the electronic product.
func (e *AudioVideoBuilder) SetRAM(ram string) {}

// SetStorageType sets the storage type of the electronic product.
func (e *AudioVideoBuilder) SetStorageType(storageType string) {}

// SetOS sets the operating system of the electronic product.
func (e *AudioVideoBuilder) SetOS(os string) {}

// SetStorageCapacity sets the storage capacity of the electronic product.
func (e *AudioVideoBuilder) SetStorageCapacity(storageCapacity string) {}
