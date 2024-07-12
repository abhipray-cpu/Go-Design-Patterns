package electronics

// Electronic represents an electronic product.
// Electronic represents an electronic product.
type Electronic struct {
	ID          string  // Unique identifier for the electronic product.
	Vendor      string  // Vendor of the electronic product.
	Name        string  // Name of the electronic product.
	Brand       string  // Brand of the electronic product.
	Price       float64 // Price of the electronic product.
	Description string  // Description of the electronic product.

	// Category specific attributes
	Category        string   // Category of the electronic product.
	ScreenSize      string   // Screen size of the electronic product.
	BatteryLife     string   // Battery life of the electronic product.
	OS              string   // Operating system of the electronic product.
	StorageCapacity string   // Storage capacity of the electronic product.
	CPU             string   // CPU (Central Processing Unit) of the electronic product.
	RAM             string   // RAM (Random Access Memory) of the electronic product.
	StorageType     string   // Storage type of the electronic product.
	Resolution      string   // Resolution of the electronic product.
	SensorType      string   // Sensor type of the electronic product.
	RefreshRate     string   // Refresh rate of the electronic product.
	OpticalZoom     string   // Optical zoom of the electronic product.
	VideoQuality    string   // Video quality of the electronic product.
	Type            string   // Type of the electronic product.
	EnergyRating    string   // Energy rating of the electronic product.
	Capacity        string   // Capacity of the electronic product.
	Connectivity    string   // Connectivity options of the electronic product.
	PowerOutput     string   // Power output of the electronic product.
	Weight          string   // Weight of the electronic product.
	Dimensions      string   // Dimensions of the electronic product.
	Features        []string // Additional features of the electronic product.
}
