package electronics

// Director is responsible for constructing electronic products using a builder.
type Director struct {
	builder ElectronicBuilder
}

// NewDirector creates a new Director instance with the given builder.
func NewDirector(builder ElectronicBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// SetBuilder sets the builder for the Director.
func (d *Director) SetBuilder(builder ElectronicBuilder) {
	d.builder = builder
}

// BuildMobile constructs a mobile product using the builder and returns the builder.
func (d *Director) BuildMobile(id, name string, price float64, description string, vendor string, brand, category, screenSize, batteryLife, os, storage, cpu, ram, resolution, prodType, weight, dimension string, features []string) ElectronicBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetScreenSize(screenSize)
	d.builder.SetBatteryLife(batteryLife)
	d.builder.SetOS(os)
	d.builder.SetStorageCapacity(storage)
	d.builder.SetCPU(cpu)
	d.builder.SetRAM(ram)
	d.builder.SetResolution(resolution)
	d.builder.SetType(prodType)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetFeatures(features)
	return d.builder
}

// BuildPC constructs a PC product using the builder and returns the builder.
func (d *Director) BuildPC(id, name string, price float64, description string, vendor string, brand, category, screenSize, batteryLife, os, storageCapacity, storageType, cpu, ram, resolution, prodType, weight, dimension string, features []string) ElectronicBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetScreenSize(screenSize)
	d.builder.SetBatteryLife(batteryLife)
	d.builder.SetOS(os)
	d.builder.SetStorageCapacity(storageCapacity)
	d.builder.SetStorageType(storageType)
	d.builder.SetCPU(cpu)
	d.builder.SetRAM(ram)
	d.builder.SetResolution(resolution)
	d.builder.SetType(prodType)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetFeatures(features)
	return d.builder
}

// BuildCamera constructs a camera product using the builder and returns the builder.
func (d *Director) BuildCamera(id, name string, price float64, description, vendor string, brand, category, screenSize string, batteryLife, storage string, resolution, sensor, opticalZoom, prodType, refreshRate, videoQuality, weight, dimension string, features []string) ElectronicBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetScreenSize(screenSize)
	d.builder.SetBatteryLife(batteryLife)
	d.builder.SetStorageCapacity(storage)
	d.builder.SetResolution(resolution)
	d.builder.SetSensorType(sensor)
	d.builder.SetOpticalZoom(opticalZoom)
	d.builder.SetType(prodType)
	d.builder.SetRefreshRate(refreshRate)
	d.builder.SetVideoQuality(videoQuality)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetFeatures(features)
	return d.builder
}

// BuildHomeAppliance constructs a home appliance product using the builder and returns the builder.
func (d *Director) BuildHomeAppliance(id, name string, price float64, description string, vendor string, brand, category string, prodType, rating, capacity, weight, dimension string, features []string) ElectronicBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetType(prodType)
	d.builder.SetEnergyRating(rating)
	d.builder.SetCapacity(capacity)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetFeatures(features)
	return d.builder
}

// BuildAudioVideo constructs an audio/video product using the builder and returns the builder.
func (d *Director) BuildAudioVideo(id, name string, price float64, description, vendor string, brand, category string, batteryLife, prodType, connectivity, power, weight, dimension string, features []string) ElectronicBuilder {
	d.builder.SetAttributes(id, name, price, description, vendor)
	d.builder.SetBrand(brand)
	d.builder.SetCategory(category)
	d.builder.SetBatteryLife(batteryLife)
	d.builder.SetType(prodType)
	d.builder.SetConnectivity(connectivity)
	d.builder.SetPowerOutput(power)
	d.builder.SetWeight(weight)
	d.builder.SetDimensions(dimension)
	d.builder.SetFeatures(features)
	return d.builder
}
