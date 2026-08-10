func rotate(nums []int, k int) {


	k = k % len(nums)

	// rotate
	left := 0
	right := len(nums)-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	fmt.Println(nums)

	// rotate from 0 to len(nums)-k
	left = 0
	right = k-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	fmt.Println(nums)

	// rotate from k+1 to n
	left = k
	right = len(nums)-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	fmt.Println(nums)
}
