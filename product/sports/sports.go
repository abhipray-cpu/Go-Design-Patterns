package sports

// Sports represents a sports product.
type Sports struct {
	ID          string // ID of the sports product.
	Name        string  // Name of the sports product.
	Category    string  // Category of the sports product.
	Price       float64 // Price of the sports product.
	Brand       string  // Brand of the sports product.
	Description string  // Description of the sports product.
	Weight      string // Weight of the sports product.
	Type		string  // Type of the sports product.
	Size 		string // size of the product
	Color 		string // color of the product
	Material    string  // Material used in the sports product.
	Dimensions  string  // Dimensions of the sports product.
	MinAge      int     // Minimum age requirement for the sports product.
	MaxAge      int     // Maximum age requirement for the sports product.
	Durability  string  // Durability level of the sports product.
	Suitable 	[]string // Suitable for the audience.
	Vendor      string  // Vendor of the sports product.
}
