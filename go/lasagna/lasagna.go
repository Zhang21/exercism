package lasagna

// OvenTime is a constant
// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven > OvenTime {
		panic("RemainingOvenTime not implemented")
	}
	remainingOvenTime := OvenTime - actualMinutesInOven
	return remainingOvenTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	preparationTime := numberOfLayers * 2
	return preparationTime
	// panic("PreparationTime not implemented")
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	elapsedTime := numberOfLayers*2 + actualMinutesInOven
	return elapsedTime
	// panic("ElapsedTime not implemented")
}
