package fashion

// Fashion represents a fashion product.
type Fashion struct {
	Id          string  // Unique identifier of the fashion product.
	Vendor 		string  // Vendor of the fashion product.
	Name        string  // Name of the fashion product.
	Price       float64 // Price of the fashion product.
	Description string  // Description of the fashion product.
	Brand       string  // Brand of the fashion product.
	Material    string  // Material used in the fashion product.

	// Category-specific attributes
	Category		string // Category of the fashion product.
	Size            string // Size of the clothing, footwear, or bags.
	Color           string // Color of the clothing, footwear, bags, or watches.
	Gender          string // Gender for which the fashion product is suitable.
	Style           string // Style of the fashion product.
	Type            string // Type of the fashion product.
	GemStone        string // Gemstone used in the fashion product.
	WaterResistance string // Water resistance level of the fashion product.
}
