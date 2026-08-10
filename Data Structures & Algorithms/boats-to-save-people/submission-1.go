func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	numberOfBoat := 0
	
	left := 0
	right := len(people)-1

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		
		numberOfBoat += 1
	}
	return numberOfBoat

}
