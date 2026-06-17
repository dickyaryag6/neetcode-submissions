func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 0

	if len(s) == 0 {
		return 0
	}

	maxLength := 1

	hashMap := map[byte]bool{}

	for {
		if l == r && l == len(s)-1 && r == len(s)-1 {
			break
		}

		if _, ok := hashMap[s[r]]; ok {
			delete(hashMap, s[l])
			l+=1
			continue
		}

		hashMap[s[r]] = true

		length := (r - l) + 1
		maxLength = max(maxLength, length)

		if r < len(s)-1 {
			r+=1
		}
		
	}

	return maxLength
}
