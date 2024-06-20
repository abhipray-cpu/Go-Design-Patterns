package electronics

// HomeApplianceBuilder is a builder for creating home appliances.
type HomeApplianceBuilder struct {
	electronic Electronic
}

// newHomeApplianceBuilder creates a new instance of HomeApplianceBuilder.
func newHomeApplianceBuilder() *HomeApplianceBuilder {
	return &HomeApplianceBuilder{}
}

// SetAttributes sets the attributes of the home appliance.
func (e *HomeApplianceBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	e.electronic.ID = id
	e.electronic.Name = name
	e.electronic.Price = price
	e.electronic.Description = description
	e.electronic.Vendor = vendor
}

// SetBrand sets the brand of the home appliance.
func (e *HomeApplianceBuilder) SetBrand(brand string) {
	e.electronic.Brand = brand
}

// SetCategory sets the category of the home appliance.
func (e *HomeApplianceBuilder) SetCategory(category string) {
	e.electronic.Category = category
}

// SetType sets the type of the home appliance.
func (e *HomeApplianceBuilder) SetType(deviceType string) {
	e.electronic.Type = deviceType
}

// SetEnergyRating sets the energy rating of the home appliance.
func (e *HomeApplianceBuilder) SetEnergyRating(energyRating string) {
	e.electronic.EnergyRating = energyRating
}

// SetCapacity sets the capacity of the home appliance.
func (e *HomeApplianceBuilder) SetCapacity(capacity string) {
	e.electronic.Capacity = capacity
}

// SetWeight sets the weight of the home appliance.
func (e *HomeApplianceBuilder) SetWeight(weight string) {
	e.electronic.Weight = weight
}

// SetDimensions sets the dimensions of the home appliance.
func (e *HomeApplianceBuilder) SetDimensions(dimensions string) {
	e.electronic.Dimensions = dimensions
}

// SetFeatures sets the features of the home appliance.
func (e *HomeApplianceBuilder) SetFeatures(features []string) {
	e.electronic.Features = features
}

// GetElectronic returns the built Electronic object.
func (e *HomeApplianceBuilder) GetElectronic() Electronic {
	return e.electronic
}

// unused exposed functions

// SetStorageType sets the storage type of the home appliance.
func (e *HomeApplianceBuilder) SetStorageType(storageType string) {
	// This function is currently unused.
}

// SetResolution sets the resolution of the home appliance.
func (e *HomeApplianceBuilder) SetResolution(resolution string) {
	// This function is currently unused.
}

// SetSensorType sets the sensor type of the home appliance.
func (e *HomeApplianceBuilder) SetSensorType(sensorType string) {
	// This function is currently unused.
}

// SetRefreshRate sets the refresh rate of the home appliance.
func (e *HomeApplianceBuilder) SetRefreshRate(refreshRate string) {
	// This function is currently unused.
}

// SetOpticalZoom sets the optical zoom of the home appliance.
func (e *HomeApplianceBuilder) SetOpticalZoom(opticalZoom string) {
	// This function is currently unused.
}

// SetVideoQuality sets the video quality of the home appliance.
func (e *HomeApplianceBuilder) SetVideoQuality(videoQuality string) {
	// This function is currently unused.
}

// SetConnectivity sets the connectivity of the home appliance.
func (e *HomeApplianceBuilder) SetConnectivity(connectivity string) {
	// This function is currently unused.
}

// SetPowerOutput sets the power output of the home appliance.
func (e *HomeApplianceBuilder) SetPowerOutput(powerOutput string) {
	// This function is currently unused.
}

// SetScreenSize sets the screen size of the home appliance.
func (e *HomeApplianceBuilder) SetScreenSize(screenSize string) {
	// This function is currently unused.
}

// SetResolutionType sets the resolution type of the home appliance.
func (e *HomeApplianceBuilder) SetResolutionType(resolution string) {
	// This function is currently unused.
}

// SetRefreshRateType sets the refresh rate type of the home appliance.
func (e *HomeApplianceBuilder) SetRefreshRateType(refreshRate string) {
	// This function is currently unused.
}

// SetOpticalZoomType sets the optical zoom type of the home appliance.
func (e *HomeApplianceBuilder) SetOpticalZoomType(opticalZoom string) {
	// This function is currently unused.
}

// SetVideoQualityType sets the video quality type of the home appliance.
func (e *HomeApplianceBuilder) SetVideoQualityType(videoQuality string) {
	// This function is currently unused.
}

// SetBatteryLife sets the battery life of the home appliance.
func (e *HomeApplianceBuilder) SetBatteryLife(batteryLife string) {
	// This function is currently unused.
}

// SetStorageCapacity sets the storage capacity of the home appliance.
func (e *HomeApplianceBuilder) SetStorageCapacity(storageCapacity string) {
	// This function is currently unused.
}

// SetCPU sets the CPU of the home appliance.
func (e *HomeApplianceBuilder) SetCPU(cpu string) {
	// This function is currently unused.
}

// SetRAM sets the RAM of the home appliance.
func (e *HomeApplianceBuilder) SetRAM(ram string) {
	// This function is currently unused.
}

// SetOS sets the operating system of the home appliance.
func (e *HomeApplianceBuilder) SetOS(os string) {
	// This function is currently unused.
}
