func checkInclusion(s1 string, s2 string) bool {

	hashMap := map[rune]int{}
	for _, char := range s1 {

		if _, ok := hashMap[char]; ok {
			hashMap[char]+=1
		} else {
			hashMap[char]=1
		}
	}

	left := 0
	right := len(s1)-1

	for {
		if right == len(s2) {
			break
		}
		subString := s2[left:right+1]
		hashMap2 := map[rune]int{}
		for _, char := range subString {
			if _, ok := hashMap2[char]; ok {
				hashMap2[char]+=1
			} else {
				hashMap2[char]=1
			}
		}
		if compare(hashMap, hashMap2) {
			return true
		}
		left+=1
		right+=1
		
	}

	return false

}

func compare(hashMap map[rune]int, hashMap2 map[rune]int) bool {
	for key, value := range hashMap {
		if val, ok := hashMap2[key]; ok {
			if value != val {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
