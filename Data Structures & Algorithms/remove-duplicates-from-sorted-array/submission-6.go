func removeDuplicates(nums []int) int {
	pointer1 := 0
	pointer2 := 1
	duplicate := false

	if len(nums) == 1 {
		return 1
	}

	for {
		if pointer1 == len(nums)-1 || pointer2 == len(nums)-1  {
			if nums[pointer1] == nums[pointer2] {
				nums = append(nums[:pointer1], nums[pointer2:]...)
			} else {
				nums = append(nums[:pointer1+1], nums[pointer2:]...)
			}
			return len(nums)
		}

		if nums[pointer1] == nums[pointer2] {
			duplicate = true
			if pointer2 < len(nums) {
				pointer2 += 1
			}
			continue
		} else if duplicate == true {

			fmt.Println(pointer1, " ", pointer2)

			fmt.Println(nums[:pointer1+1]," ", nums[pointer2:])

			nums = append(nums[:pointer1+1], nums[pointer2:]...)
			duplicate = false

			pointer2 = pointer1 + 1
 
		} else {
			pointer1 += 1
			pointer2 += 1
		}
	}

	return len(nums)
}
