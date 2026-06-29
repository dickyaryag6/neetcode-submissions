func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums)*2)

	for index, _ := range ans {
		if index < len(nums) {
			ans[index] = nums[index]
		} else {
			ans[index] = nums[index-len(nums)]
		}
	}

	return ans
}
