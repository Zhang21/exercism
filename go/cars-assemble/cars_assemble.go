package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
// Attention: return float64
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	carsH := float64(float64(productionRate) * successRate / 100)
	return carsH
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
// Attention: return int
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsM := int(float64(productionRate) * successRate / 100 / float64(60))
	return carsM
}

// CalculateCost works out the cost of producing the given number of cars
// Attention: return uint
func CalculateCost(carsCount int) uint {
	x := carsCount / 10
	y := carsCount % 10
	costs := uint(x*95000 + y*10000)
	return costs
}
