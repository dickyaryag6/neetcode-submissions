func maxProfit(prices []int) int {
	// find the lowest price with the earliest day
	// find the highest price

	min := 101

	max := 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if (price - min) > max {
			max = price-min
		}
	}

	return max

}
