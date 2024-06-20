// Package electronics provides functions to create different types of electronic products.
package electronics

// createMobileProduct creates a mobile product using the given parameters and returns an Electronic object.
// createMobileProduct creates a new mobile product with the given parameters and returns an Electronic object.
func createMobileProduct(id, name string, price float64, description string, vendor string, brand, category, screenSize, batteryLife, os, storage, cpu, ram, resolution, prodType, weight, dimension string, features []string) Electronic {
	builder := getBuilder("mobile")
	director := NewDirector(builder)
	build := director.BuildMobile(id, name, price, description, vendor, brand, category, screenSize, batteryLife, os, storage, cpu, ram, resolution, prodType, weight, dimension, features)
	return build.GetElectronic()
}

// createPCProduct creates a PC product using the given parameters and returns an Electronic object.
func createPCProduct(id, name string, price float64, description string, vendor string, brand, category, screenSize, batteryLife, os, storageCapacity, storageType, cpu, ram, resolution, prodType, weight, dimension string, features []string) Electronic {
	builder := getBuilder("pc")
	director := NewDirector(builder)
	build := director.BuildPC(id, name, price, description, vendor, brand, category, screenSize, batteryLife, os, storageCapacity, storageType, cpu, ram, resolution, prodType, weight, dimension, features)
	return build.GetElectronic()
}

// createCameraProduct creates a camera product using the given parameters and returns an Electronic object.
func createCameraProduct(id, name string, price float64, description, vendor string, brand, category, screenSize string, batteryLife, storage string, resolution, sensor, opticalZoom, prodType, refreshRate, videoQuality, weight, dimension string, features []string) Electronic {
	builder := getBuilder("camera")
	director := NewDirector(builder)
	build := director.BuildCamera(id, name, price, description, vendor, brand, category, screenSize, batteryLife, storage, resolution, sensor, opticalZoom, prodType, refreshRate, videoQuality, weight, dimension, features)
	return build.GetElectronic()
}

// createAudioVideoProduct creates an audio/video product using the given parameters and returns an Electronic object.
func createAudioVideoProduct(id, name string, price float64, description, vendor string, brand, category string, batteryLife, prodType, connectivity, power, weight, dimension string, features []string) Electronic {
	builder := getBuilder("audio")
	director := NewDirector(builder)
	build := director.BuildAudioVideo(id, name, price, description, vendor, brand, category, batteryLife, prodType, connectivity, power, weight, dimension, features)
	return build.GetElectronic()
}

// createHomeApplianceProduct creates a home appliance product using the given parameters and returns an Electronic object.
func createHomeApplianceProduct(id, name string, price float64, description string, vendor string, brand, category string, prodType, rating, capacity, weight, dimension string, features []string) Electronic {
	builder := getBuilder("appliance")
	director := NewDirector(builder)
	build := director.BuildHomeAppliance(id, name, price, description, vendor, brand, category, prodType, rating, capacity, weight, dimension, features)
	return build.GetElectronic()
}
