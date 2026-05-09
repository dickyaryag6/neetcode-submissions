func isAnagram(s string, t string) bool {
	hashMap := map[byte]int{}
	hashMap2 := map[byte]int{}

	if len(s) != len(t) {
		return false
	}

	for i:=0; i<len(s); i++ {
		if _, ok := hashMap[s[i]]; ok {
			hashMap[s[i]] = hashMap[s[i]]+1
		} else {
			hashMap[s[i]] = 1
		}
	}

	for i:=0; i<len(t); i++ {
		if _, ok := hashMap2[t[i]]; ok {
			hashMap2[t[i]] = hashMap2[t[i]]+1
		} else {
			hashMap2[t[i]] = 1
		}
	}

	for key, value := range hashMap {
		if value2, ok := hashMap2[key]; !ok {
			return false
		} else {
			if value != value2 {
				return false
			}
		}
	}

	return true

}
