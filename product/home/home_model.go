package home

// Home represents a product that belongs to the "Home" category.
type Home struct {
	ID          string  // Unique identifier for the home product.
	Vendor      string  // Vendor of the home product.
	Name        string  // Name of the home product.
	Description string  // Description of the home product.
	Category    string  // Category of the home product.
	Price       float64 // Price of the home product.
	Brand       string  // Brand of the home product.
	Material    string  // Material used in the home product.
	Dimensions  string  // Dimensions of the home product.
	Weight      string  // Weight of the home product.

	// Category specific attributes
	Type             string // Type of the home product.
	PowerSource      string // Power source required for the home product.
	AssemblyRequired bool   // Indicates if assembly is required for the home product.
	Volume           string // Volume of the home product.
	Scent            string // Scent of the home product.
	Color            string // Color of the home product.
	Style            string // Style of the home product.
	Capacity         string // Capacity of the home product.
	DishWasherSafe   bool   // Indicates if the home product is dishwasher safe.
}
