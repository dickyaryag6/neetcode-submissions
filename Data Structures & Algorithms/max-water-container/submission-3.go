func maxArea(heights []int) int {


	p1 := 0
	p2 := len(heights)-1

	result := 0

	for {
		if p1 == p2 {
			break
		}

		curr := (p2-p1) * min(heights[p1], heights[p2])
		if curr > result {
			result = curr
		}

		if heights[p1] == min(heights[p1], heights[p2]) {
			p1+=1
		} else {
			p2-=1
		}

	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
