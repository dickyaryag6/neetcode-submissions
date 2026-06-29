func replaceElements(arr []int) []int {
	biggest := arr[len(arr)-1]
	for i:=len(arr)-1;i>=0;i-- {
		current := arr[i]
		arr[i] = biggest

		if current > biggest {
			biggest = current
		}
	}

	arr[len(arr)-1] = -1

	return arr
}
