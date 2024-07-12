package electronics

// PCBuilder is a builder for creating PC objects.
type PCBuilder struct {
	electronic Electronic
}

// newPCBuilder creates a new instance of PCBuilder.
func newPCBuilder() *PCBuilder {
	return &PCBuilder{}
}

// SetAttributes sets the attributes of the PC.
func (e *PCBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	e.electronic.ID = id
	e.electronic.Name = name
	e.electronic.Price = price
	e.electronic.Description = description
	e.electronic.Vendor = vendor
}

// SetBrand sets the brand of the PC.
func (e *PCBuilder) SetBrand(brand string) {
	e.electronic.Brand = brand
}

// SetCategory sets the category of the PC.
func (e *PCBuilder) SetCategory(category string) {
	e.electronic.Category = category
}

// SetOS sets the operating system of the PC.
func (e *PCBuilder) SetOS(os string) {
	e.electronic.OS = os
}

// SetScreenSize sets the screen size of the PC.
func (e *PCBuilder) SetScreenSize(screenSize string) {
	e.electronic.ScreenSize = screenSize
}

// SetBatteryLife sets the battery life of the PC.
func (e *PCBuilder) SetBatteryLife(batteryLife string) {
	e.electronic.BatteryLife = batteryLife
}

// SetStorageType sets the storage type of the PC.
func (e *PCBuilder) SetStorageType(storageType string) {
	e.electronic.StorageType = storageType
}

// SetStorageCapacity sets the storage capacity of the PC.
func (e *PCBuilder) SetStorageCapacity(storageCapacity string) {
	e.electronic.StorageCapacity = storageCapacity
}

// SetCPU sets the CPU of the PC.
func (e *PCBuilder) SetCPU(cpu string) {
	e.electronic.CPU = cpu
}

// SetRAM sets the RAM of the PC.
func (e *PCBuilder) SetRAM(ram string) {
	e.electronic.RAM = ram
}

// SetResolution sets the resolution of the PC.
func (e *PCBuilder) SetResolution(resolution string) {
	e.electronic.Resolution = resolution
}

// SetType sets the type of the PC.
func (e *PCBuilder) SetType(deviceType string) {
	e.electronic.Type = deviceType
}

// SetWeight sets the weight of the PC.
func (e *PCBuilder) SetWeight(weight string) {
	e.electronic.Weight = weight
}

// SetDimensions sets the dimensions of the PC.
func (e *PCBuilder) SetDimensions(dimensions string) {
	e.electronic.Dimensions = dimensions
}

// SetFeatures sets the features of the PC.
func (e *PCBuilder) SetFeatures(features []string) {
	e.electronic.Features = features
}

// GetElectronic returns the built PC object.
func (e *PCBuilder) GetElectronic() Electronic {
	return e.electronic
}

// The following functions are unused and exposed.

// SetSensorType sets the sensor type of the PC.
func (e *PCBuilder) SetSensorType(sensorType string) {}

// SetRefreshRate sets the refresh rate of the PC.
func (e *PCBuilder) SetRefreshRate(refreshRate string) {}

// SetOpticalZoom sets the optical zoom of the PC.
func (e *PCBuilder) SetOpticalZoom(opticalZoom string) {}

// SetVideoQuality sets the video quality of the PC.
func (e *PCBuilder) SetVideoQuality(videoQuality string) {}

// SetEnergyRating sets the energy rating of the PC.
func (e *PCBuilder) SetEnergyRating(energyRating string) {}

// SetCapacity sets the capacity of the PC.
func (e *PCBuilder) SetCapacity(capacity string) {}

// SetConnectivity sets the connectivity of the PC.
func (e *PCBuilder) SetConnectivity(connectivity string) {}

// SetPowerOutput sets the power output of the PC.
func (e *PCBuilder) SetPowerOutput(powerOutput string) {}
