package sports

// ExerciseBuilder is a builder for creating Exercise objects.
type ExerciseBuilder struct {
	sports Sports
}

// newExerciseBuilder creates a new instance of ExerciseBuilder.
func newExerciseBuilder() *ExerciseBuilder {
	return &ExerciseBuilder{}
}

// SetAttributes sets the attributes of the Exercise.
func (e *ExerciseBuilder) SetAttributes(id string, name string, price float64, description string) {
	e.sports.ID = id
	e.sports.Name = name
	e.sports.Price = price
	e.sports.Description = description
}

// SetCategory sets the category of the Exercise.
func (e *ExerciseBuilder) SetCategory(category string) {
	e.sports.Category = category
}

// SetAgeGroup sets the age group of the Exercise.
func (e *ExerciseBuilder) SetAgeGroup(minAge, maxAge int) {
	e.sports.MinAge = minAge
	e.sports.MaxAge = maxAge
}

// SetWeight sets the weight of the Exercise.
func (e *ExerciseBuilder) SetWeight(weight string) {
	e.sports.Weight = weight
}

// SetMaterial sets the material of the Exercise.
func (e *ExerciseBuilder) SetMaterial(material string) {
	e.sports.Material = material
}

// SetDimensions sets the dimensions of the Exercise.
func (e *ExerciseBuilder) SetDimensions(dimension string) {
	e.sports.Dimensions = dimension
}

// SetDurability sets the durability of the Exercise.
func (e *ExerciseBuilder) SetDurability(durability string) {
	e.sports.Durability = durability
}

// SetVendor sets the vendor of the Exercise.
func (e *ExerciseBuilder) SetVendor(vendor string) {
	e.sports.Vendor = vendor
}

// SetType sets the type of the Exercise.
func (e *ExerciseBuilder) SetType(prodType string) {
	e.sports.Type = prodType
}

// SetSuitableFor sets the suitableFor field of the Exercise.
func (e *ExerciseBuilder) SetSuitableFor(suitableFor []string) {
	e.sports.Suitable = suitableFor
}

// SetSize sets the size of the Exercise.
func (e *ExerciseBuilder) SetSize(size string) {
	e.sports.Size = size
}

// SetColor sets the color of the Exercise.
func (e *ExerciseBuilder) SetColor(color string) {
	e.sports.Color = color
}

// SetBrand sets the brand of the Exercise.
func (e *ExerciseBuilder) SetBrand(brand string) {
	e.sports.Brand = brand
}

// GetItem returns the built Exercise object.
func (e *ExerciseBuilder) GetItem() Sports {
	return e.sports
}

// unused functions from abstract interface
