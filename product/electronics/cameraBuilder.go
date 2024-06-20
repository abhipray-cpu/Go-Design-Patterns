package electronics

// CameraBuilder is a builder for creating Camera objects.
type CameraBuilder struct {
	electronic Electronic
}

// newCameraBuilder creates a new instance of CameraBuilder.
func newCameraBuilder() *CameraBuilder {
	return &CameraBuilder{}
}

// SetAttributes sets the attributes of the Camera.
func (e *CameraBuilder) SetAttributes(id, name string, price float64, description, vendor string) {
	e.electronic.ID = id
	e.electronic.Name = name
	e.electronic.Price = price
	e.electronic.Description = description
	e.electronic.Vendor = vendor
}

// SetBrand sets the brand of the Camera.
func (e *CameraBuilder) SetBrand(brand string) {
	e.electronic.Brand = brand
}

// SetCategory sets the category of the Camera.
func (e *CameraBuilder) SetCategory(category string) {
	e.electronic.Category = category
}

// SetScreenSize sets the screen size of the Camera.
func (e *CameraBuilder) SetScreenSize(screenSize string) {
	e.electronic.ScreenSize = screenSize
}

// SetBatteryLife sets the battery life of the Camera.
func (e *CameraBuilder) SetBatteryLife(batteryLife string) {
	e.electronic.BatteryLife = batteryLife
}

// SetStorageCapacity sets the storage capacity of the Camera.
func (e *CameraBuilder) SetStorageCapacity(storageCapacity string) {
	e.electronic.StorageCapacity = storageCapacity
}

// SetResolution sets the resolution of the Camera.
func (e *CameraBuilder) SetResolution(resolution string) {
	e.electronic.Resolution = resolution
}

// SetSensorType sets the sensor type of the Camera.
func (e *CameraBuilder) SetSensorType(sensorType string) {
	e.electronic.SensorType = sensorType
}

// SetOpticalZoom sets the optical zoom of the Camera.
func (e *CameraBuilder) SetOpticalZoom(opticalZoom string) {
	e.electronic.OpticalZoom = opticalZoom
}

// SetType sets the device type of the Camera.
func (e *CameraBuilder) SetType(deviceType string) {
	e.electronic.Type = deviceType
}

// SetRefreshRate sets the refresh rate of the Camera.
func (e *CameraBuilder) SetRefreshRate(refreshRate string) {
	e.electronic.RefreshRate = refreshRate
}

// SetVideoQuality sets the video quality of the Camera.
func (e *CameraBuilder) SetVideoQuality(videoQuality string) {
	e.electronic.VideoQuality = videoQuality
}

// SetWeight sets the weight of the Camera.
func (e *CameraBuilder) SetWeight(weight string) {
	e.electronic.Weight = weight
}

// SetDimensions sets the dimensions of the Camera.
func (e *CameraBuilder) SetDimensions(dimensions string) {
	e.electronic.Dimensions = dimensions
}

// SetFeatures sets the features of the Camera.
func (e *CameraBuilder) SetFeatures(features []string) {
	e.electronic.Features = features
}

// GetElectronic returns the built Camera object.
func (e *CameraBuilder) GetElectronic() Electronic {
	return e.electronic
}

// The following functions are unused and exposed, but not implemented.

// SetEnergyRating sets the energy rating of the Camera.
func (e *CameraBuilder) SetEnergyRating(energyRating string) {
	// Not implemented
}

// SetCapacity sets the capacity of the Camera.
func (e *CameraBuilder) SetCapacity(capacity string) {
	// Not implemented
}

// SetConnectivity sets the connectivity of the Camera.
func (e *CameraBuilder) SetConnectivity(connectivity string) {
	// Not implemented
}

// SetPowerOutput sets the power output of the Camera.
func (e *CameraBuilder) SetPowerOutput(powerOutput string) {
	// Not implemented
}

// SetStorageType sets the storage type of the Camera.
func (e *CameraBuilder) SetStorageType(storageType string) {
	// Not implemented
}

// SetRAM sets the RAM of the Camera.
func (e *CameraBuilder) SetRAM(ram string) {
	// Not implemented
}

// SetCPU sets the CPU of the Camera.
func (e *CameraBuilder) SetCPU(cpu string) {
	// Not implemented
}

// SetOS sets the operating system of the Camera.
func (e *CameraBuilder) SetOS(os string) {
	// Not implemented
}
