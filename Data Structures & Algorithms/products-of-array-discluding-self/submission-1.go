func productExceptSelf(nums []int) []int {
	totalProduct := 1
	zeroExist := false
	multipleZeros := false
	for _, num := range nums {
		if zeroExist && num == 0 {
			multipleZeros = true
		} else if num == 0 {
			zeroExist = true
		} else {
			totalProduct *= num
		}
	}

	if multipleZeros {
		for i := range nums {
			nums[i] = 0
		}
		return nums
	}

	for i, num := range nums {
		if num == 0 {
			nums[i] = totalProduct
		} else if zeroExist {
			nums[i] = 0
		} else {
			nums[i] = totalProduct/num
		}
	}

	return nums
}
