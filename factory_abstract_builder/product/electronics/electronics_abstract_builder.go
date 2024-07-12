package electronics

// ElectronicBuilder is an interface that defines the methods for building an electronic product.
type ElectronicBuilder interface {
	// SetAttributes sets the attributes of the electronic product.
	// It takes the ID, name, price, description, and vendor as parameters.
	SetAttributes(id, name string, price float64, description, vendor string)

	// SetBrand sets the brand of the electronic product.
	// It takes the brand as a parameter.
	SetBrand(brand string)

	// SetCategory sets the category of the electronic product.
	// It takes the category as a parameter.
	SetCategory(category string)

	// SetScreenSize sets the screen size of the electronic product.
	// It takes the screen size as a parameter.
	SetScreenSize(screenSize string)

	// SetBatteryLife sets the battery life of the electronic product.
	// It takes the battery life as a parameter.
	SetBatteryLife(batteryLife string)

	// SetOS sets the operating system of the electronic product.
	// It takes the operating system as a parameter.
	SetOS(os string)

	// SetStorageCapacity sets the storage capacity of the electronic product.
	// It takes the storage capacity as a parameter.
	SetStorageCapacity(storageCapacity string)

	// SetCPU sets the CPU of the electronic product.
	// It takes the CPU as a parameter.
	SetCPU(cpu string)

	// SetRAM sets the RAM of the electronic product.
	// It takes the RAM as a parameter.
	SetRAM(ram string)

	// SetStorageType sets the storage type of the electronic product.
	// It takes the storage type as a parameter.
	SetStorageType(storageType string)

	// SetResolution sets the resolution of the electronic product.
	// It takes the resolution as a parameter.
	SetResolution(resolution string)

	// SetSensorType sets the sensor type of the electronic product.
	// It takes the sensor type as a parameter.
	SetSensorType(sensorType string)

	// SetRefreshRate sets the refresh rate of the electronic product.
	// It takes the refresh rate as a parameter.
	SetRefreshRate(refreshRate string)

	// SetOpticalZoom sets the optical zoom of the electronic product.
	// It takes the optical zoom as a parameter.
	SetOpticalZoom(opticalZoom string)

	// SetVideoQuality sets the video quality of the electronic product.
	// It takes the video quality as a parameter.
	SetVideoQuality(videoQuality string)

	// SetType sets the type of the electronic product.
	// It takes the type as a parameter.
	SetType(eType string)

	// SetEnergyRating sets the energy rating of the electronic product.
	// It takes the energy rating as a parameter.
	SetEnergyRating(energyRating string)

	// SetCapacity sets the capacity of the electronic product.
	// It takes the capacity as a parameter.
	SetCapacity(capacity string)

	// SetConnectivity sets the connectivity options of the electronic product.
	// It takes the connectivity options as a parameter.
	SetConnectivity(connectivity string)

	// SetPowerOutput sets the power output of the electronic product.
	// It takes the power output as a parameter.
	SetPowerOutput(powerOutput string)

	// SetWeight sets the weight of the electronic product.
	// It takes the weight as a parameter.
	SetWeight(weight string)

	// SetDimensions sets the dimensions of the electronic product.
	// It takes the dimensions as a parameter.
	SetDimensions(dimensions string)

	// SetFeatures sets the features of the electronic product.
	// It takes the features as a parameter.
	SetFeatures(features []string)

	// GetElectronic returns the built electronic product.
	GetElectronic() Electronic
}

// getBuilder returns an instance of ElectronicBuilder based on the provided builderType.
// If builderType is "mobile", it returns a new instance of MobileBuilder.
// If builderType is "pc", it returns a new instance of PCBuilder.
// If builderType is "camera", it returns a new instance of CameraBuilder.
// If builderType is "audio", it returns a new instance of AudioVideoBuilder.
// If builderType is "appliance", it returns a new instance of HomeApplianceBuilder.
// If builderType is not one of the above, it returns nil.
func getBuilder(builderType string) ElectronicBuilder {
	if builderType == "mobile" {
		return newMobileBuilder()
	}

	if builderType == "pc" {
		return newPCBuilder()
	}
	if builderType == "camera" {
		return newCameraBuilder()
	}
	if builderType == "audio" {
		return newAudioVideoBuilder()
	}
	if builderType == "appliance" {
		return newHomeApplianceBuilder()
	}
	return nil
}
