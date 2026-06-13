func minEatingSpeed(piles []int, h int) int {

	maxPile := 0
    for _, p := range piles {
        if p > maxPile {
            maxPile = p
        }
    }
	for k := 1; k <= maxPile; k++ {
        if canFinish(piles, k, h) {
            return k
        }
    }

    return maxPile

}

func canFinish(piles []int, k, h int) bool {
    hours := 0
    for _, p := range piles {
		// check hours to eat bananas per pile
        hours += (p + k - 1) / k  // this is ceil(p / k)
    }
    return hours <= h
}
