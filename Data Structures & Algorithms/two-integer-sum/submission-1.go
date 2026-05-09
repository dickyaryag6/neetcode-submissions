func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}

	for index, num := range nums {
		if value, ok := hashMap[num]; !ok {
			hashMap[target-num] = index
		} else {
			return []int{value, index}
		}
	}

	return []int{}
}

