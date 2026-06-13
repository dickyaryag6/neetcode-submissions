func trap(height []int) int {
	result := 0 

	n := len(height)
    if n == 0 {
        return 0
    }

	maxLeft := make([]int, n)
    maxRight := make([]int, n)

    // fill maxLeft: running max from left
    maxLeft[0] = height[0]
    for i := 1; i < n; i++ {
        maxLeft[i] = max(maxLeft[i-1], height[i])
    }

    // fill maxRight: running max from right
    maxRight[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], height[i])
    }

	
    // for each column, water = min(maxLeft, maxRight) - height
    for i := 0; i < n; i++ {
        result += min(maxLeft[i], maxRight[i]) - height[i]
    }
	return result
}
