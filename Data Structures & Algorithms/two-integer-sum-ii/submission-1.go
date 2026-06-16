func twoSum(numbers []int, target int) []int {
	p1 := 0
	p2 := len(numbers)-1

	for {
		if p1 == p2 {
			break
		}
		curr := numbers[p1]+numbers[p2]
		if curr == target {
			return []int{p1+1, p2+1}
		}

		if curr > target {
			p2-=1
		} else {
			p1+=1
		}
	}

	return []int{}
}
