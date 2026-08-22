import (
	"slices"
)

func shipWithinDays(weights []int, days int) int {
    left := slices.Max(weights)

    right := 0
    for _, weight := range weights {
        right += weight
    }

    for left < right {
        mid := left + (right - left)/2

        // check if can contain within days
        // goes left when current capacity can contain
        if canContain(weights, mid, days) {
            right = mid
            // currentCap = mid
        } else {  // goes right when it can't
            left = mid + 1
        }
    }

    return left
}

func canContain(weights []int, maxCap int, days int) bool {
    currentCargo := 0
    numOfDays := 1
    for i := len(weights)-1;i >= 0;i-- {
        if currentCargo+weights[i]<= maxCap {
            currentCargo += weights[i]
        } else {
            currentCargo = weights[i]
            numOfDays += 1
        }
    }

    return numOfDays <= days
}