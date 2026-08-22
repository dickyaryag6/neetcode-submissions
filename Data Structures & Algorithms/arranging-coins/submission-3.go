func arrangeCoins(n int) int {
	left := 0
	right := n+1

	for left < right {
		mid := left + (right-left)/2
		// coins needed in mid is enough, goes right
		if isCoinEnough(mid, n) {
			left = mid+1
		} else {
			// else left
			right = mid
		}
	}

	return left-1
}

func isCoinEnough(index, n int) bool {
	return index*(index+1)/2 <= n
}
