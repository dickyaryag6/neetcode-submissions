func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	pointerS := 0
	pointerSLength := len(s)

	for _, tChar := range t {

		if tChar == rune(s[pointerS]) {
			pointerSLength--
			pointerS++

			if pointerS > len(s) || pointerSLength == 0 {
				break
			}
		}
	}

	if pointerSLength != 0 {
		return false
	}

	return true
}
