package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	if len(birdsPerDay) == 0 {
		panic("No values in birdsPerDay!")
	}
	sum := 0
	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	if week == 0 {
		panic("Wrong week number!")
	}
	start := (week - 1) * 7
	stop := (week * 7)
	return TotalBirdCount(birdsPerDay[start:stop])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	if len(birdsPerDay) == 0 {
		panic("No values in birdsPerDay!")
	}
	// it spends 1 time
	// we do not need to exe k%2 == 1
	/*
		for k := range birdsPerDay {
			if k % 2 == 0 {
				birdsPerDay[k]++
			}
		}
	*/
	// this only spend 1/2 time
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
