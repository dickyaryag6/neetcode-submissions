func topKFrequent(nums []int, k int) []int {
	numCounts := map[int]int{}
	uniqueCount := []int{}
	for _, num := range nums {
		if _, ok := numCounts[num]; !ok {
			numCounts[num] = 1
			uniqueCount = append(uniqueCount, num)
		} else {
			numCounts[num] += 1
		}
	}

	sortedCount := sortFrequency(numCounts, uniqueCount)
	return sortedCount[len(sortedCount)-k:]
}

func sortFrequency(hashMap map[int]int, nums []int) []int {
	for i:=1; i<len(nums); i++ {
       pointer := i-1
	   current := i
       for {
           if pointer < 0 {
               break
           }

            if hashMap[nums[pointer]] > hashMap[nums[current]] {
				currentNum := nums[current]
				prevNum := nums[pointer]
				nums[current] = prevNum
				nums[pointer] = currentNum
               	current = pointer
           }
           pointer-=1
       }	
   }

	fmt.Println(nums)

   return nums
}
