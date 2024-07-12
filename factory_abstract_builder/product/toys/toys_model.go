package toys

// Toy represents a toy product.
type Toy struct {
	ID          string   // ID of the toy.
	Name        string   // Name of the toy.
	Price       float64  // Price of the toy.
	Description string   // Description of the toy.
	Category    string   // Category of the toy.
	MinAge      int      // Minimum age requirement for the toy.
	MaxAge      int      // Maximum age requirement for the toy.
	Brand       string   // Brand of the toy.
	Weight      string   // Weight of the toy.
	Material    string   // Material used in the toy.
	Color       string   // Color of the toy.
	Size        string   // size of the toy
	Suitable    []string // Suitable for the audience.
	Type        string   // Type of the toy.
	Dimensions  string   // Dimensions of the toy.
	Vendor      string   // Vendor of the toy.
}
