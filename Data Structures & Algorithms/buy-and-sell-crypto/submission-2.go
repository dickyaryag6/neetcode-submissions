func maxProfit(prices []int) int {

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
