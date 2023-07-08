package lasagna

const (
	// OvenTime is the number of minutes the lasagna should be in the oven.
	OvenTime = 40
	// layerPreparationTime is the number of minutes it takes to prepare a single layer.
	layerPreparationTime = 2
)

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * layerPreparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	timeNeeded := PreparationTime(numberOfLayers)
	return timeNeeded + actualMinutesInOven
}
