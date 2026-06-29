func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	pointerS := 0

	for _, tChar := range t {
		if pointerS >= len(s) {
			break
		}
		if tChar == rune(s[pointerS]) {
			pointerS++
		}
	}

	return pointerS == len(s)
}
