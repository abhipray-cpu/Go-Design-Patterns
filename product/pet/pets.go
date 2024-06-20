package pet

// Pet represents a product in the pet category.
// Pet represents a product in the pet category.
type Pet struct {
	ID          string  // Unique identifier for the pet
	Name        string  // Name of the pet
	Category    string  // Category of the pet
	Price       float64 // Price of the pet
	Brand       string  // Brand of the pet
	Description string  // Description of the pet
	Vendor 		string // Vendor of the pet
	// Category-Specific Attributes
	Ingredients      []string // Primarily for Food: List of ingredients
	NutritionalValue string   // Primarily for Food: Nutritional value information
	Weight           string  // For Food, some Accessories: Weight of the pet
	SuitableFor      []string   // For all categories: Suitable for which pets
	Type             string   // For Grooming, Accessories: Type of the pet
	Material         string   // For Toys, Accessories: Material of the pet
	Size             string   // For Toys, Accessories, some Grooming items: Size of the pet
	Volume           string  // For Grooming items: Volume of the pet
	Durability       string   // For Toys: Durability of the pet
}
