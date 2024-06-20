package electronics

// MobileBuilder is a builder for creating mobile electronic products.
type MobileBuilder struct {
	electronic Electronic
}

// newMobileBuilder creates a new instance of MobileBuilder.
func newMobileBuilder() *MobileBuilder {
	return &MobileBuilder{}
}

// SetAttributes sets the attributes of the mobile electronic product.
func (e *MobileBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	e.electronic.ID = id
	e.electronic.Name = name
	e.electronic.Price = price
	e.electronic.Description = description
	e.electronic.Vendor = vendor
}

// SetBrand sets the brand of the mobile electronic product.
func (e *MobileBuilder) SetBrand(brand string) {
	e.electronic.Brand = brand
}

// SetCategory sets the category of the mobile electronic product.
func (e *MobileBuilder) SetCategory(category string) {
	e.electronic.Category = category
}

// SetScreenSize sets the screen size of the mobile electronic product.
func (e *MobileBuilder) SetScreenSize(screenSize string) {
	e.electronic.ScreenSize = screenSize
}

// SetBatteryLife sets the battery life of the mobile electronic product.
func (e *MobileBuilder) SetBatteryLife(batteryLife string) {
	e.electronic.BatteryLife = batteryLife
}

// SetOS sets the operating system of the mobile electronic product.
func (e *MobileBuilder) SetOS(os string) {
	e.electronic.OS = os
}

// SetStorageCapacity sets the storage capacity of the mobile electronic product.
func (e *MobileBuilder) SetStorageCapacity(storageCapacity string) {
	e.electronic.StorageCapacity = storageCapacity
}

// SetCPU sets the CPU of the mobile electronic product.
func (e *MobileBuilder) SetCPU(cpu string) {
	e.electronic.CPU = cpu
}

// SetRAM sets the RAM of the mobile electronic product.
func (e *MobileBuilder) SetRAM(ram string) {
	e.electronic.RAM = ram
}

// SetResolution sets the resolution of the mobile electronic product.
func (e *MobileBuilder) SetResolution(resolution string) {
	e.electronic.Resolution = resolution
}

// SetType sets the type of the mobile electronic product.
func (e *MobileBuilder) SetType(deviceType string) {
	e.electronic.Type = deviceType
}

// SetWeight sets the weight of the mobile electronic product.
func (e *MobileBuilder) SetWeight(weight string) {
	e.electronic.Weight = weight
}

// SetDimensions sets the dimensions of the mobile electronic product.
func (e *MobileBuilder) SetDimensions(dimensions string) {
	e.electronic.Dimensions = dimensions
}

// SetFeatures sets the features of the mobile electronic product.
func (e *MobileBuilder) SetFeatures(features []string) {
	e.electronic.Features = features
}

// GetElectronic returns the built electronic product.
func (e *MobileBuilder) GetElectronic() Electronic {
	return e.electronic
}

// The following functions are unused and exposed, but not implemented.

// SetStorageType sets the storage type of the mobile electronic product.
func (e *MobileBuilder) SetStorageType(storageType string) {
	// Not implemented
}

// SetSensorType sets the sensor type of the mobile electronic product.
func (e *MobileBuilder) SetSensorType(sensorType string) {
	// Not implemented
}

// SetRefreshRate sets the refresh rate of the mobile electronic product.
func (e *MobileBuilder) SetRefreshRate(refreshRate string) {
	// Not implemented
}

// SetOpticalZoom sets the optical zoom of the mobile electronic product.
func (e *MobileBuilder) SetOpticalZoom(opticalZoom string) {
	// Not implemented
}

// SetVideoQuality sets the video quality of the mobile electronic product.
func (e *MobileBuilder) SetVideoQuality(videoQuality string) {
	// Not implemented
}

// SetEnergyRating sets the energy rating of the mobile electronic product.
func (e *MobileBuilder) SetEnergyRating(energyRating string) {
	// Not implemented
}

// SetCapacity sets the capacity of the mobile electronic product.
func (e *MobileBuilder) SetCapacity(capacity string) {
	// Not implemented
}

// SetConnectivity sets the connectivity of the mobile electronic product.
func (e *MobileBuilder) SetConnectivity(connectivity string) {
	// Not implemented
}

// SetPowerOutput sets the power output of the mobile electronic product.
func (e *MobileBuilder) SetPowerOutput(powerOutput string) {
	// Not implemented
}
