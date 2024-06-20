package toys

// createActionToy creates an action toy using the provided parameters.
// It uses the builder pattern to construct the toy object.
func createActionToy(id, name string, price float64, description string, category string, minAge, maxAge int, brand string, weight string, vendor string, material, color string, size, prodType, dimension string, suitable []string) Toy {
	builder := getBuilder("Action")
	director := newDirector(builder)
	toy := director.buildActionToy(id, name, price, description, category, minAge, maxAge, brand, weight, vendor, material, color, size, prodType, dimension, suitable)
	return toy.GetToy()
}

// createBoardGame creates a board game using the provided parameters.
// It uses the builder pattern to construct the toy object.
func createBoardGame(id, name string, price float64, description string, category string, minAge, maxAge int, brand string, weight string, vendor string, prodType, dimension string, suitable []string) Toy {
	builder := getBuilder("Board")
	director := newDirector(builder)
	toy := director.buildBoardToy(id, name, price, description, category, minAge, maxAge, brand, weight, vendor, prodType, dimension, suitable)
	return toy.GetToy()
}

// createVideoGame creates a video game using the provided parameters.
// It uses the builder pattern to construct the toy object.
func createVideoGame(id, name string, price float64, description string, category string, minAge, maxAge int, brand string, weight string, vendor string, suitable []string) Toy {
	builder := getBuilder("Video")
	director := newDirector(builder)
	toy := director.buildVideoToy(id, name, price, description, category, minAge, maxAge, brand, weight, vendor, suitable)
	return toy.GetToy()
}
