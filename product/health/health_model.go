package health

// Health represents a health product.
// Health represents a health product.
type Health struct {
	Id                string   // Id represents the unique identifier of the health product.
	Vendor            string   // Vendor represents the vendor of the health product.
	Name              string   // Name represents the name of the health product.
	Brand             string   // Brand represents the brand of the health product.
	Price             float64  // Price represents the price of the health product.
	Description       string   // Description represents the description of the health product.
	Category          string   // Category represents the category of the health product.
	Ingredients       []string // Ingredients represents the list of ingredients in the health product.
	UsageInstructions []string // UsageInstructions represents the usage instructions for the health product.

	// category specific attributes
	Type              string // Type represents the category type of the health product (e.g., skin type, makeup type, hair type, product type, and supplement type).
	SPF               string // SPF represents the sun protection factor of the health product.
	Volume            string // Volume represents the volume of the health product.
	Color             string // Color represents the color of the health product.
	Finish            string // Finish represents the finish of the health product.
	SulfateFree       bool   // SulfateFree indicates whether the health product is sulfate-free or not.
	ForMenWomen       string // ForMenWomen represents the target audience of the health product (e.g., for men, for women, unisex).
	SupplementForm    string // SupplementForm represents the form of the health product (e.g., tablet, capsule, powder).
	RecommendedDosage string // RecommendedDosage represents the recommended dosage of the health product.
}
