func isPalindrome(s string) bool {
	pointer1 := 0
	pointer2 := len(s)-1

	for {

		if pointer1 >= pointer2 || pointer1 > len(s)-1 || pointer2 < 0 {
			break
		}

		firstChar := s[pointer1]
		secondChar := s[pointer2]

		if !isAlphanumeric(firstChar) {
			pointer1+=1
			continue
		}

		if !isAlphanumeric(secondChar) {
			pointer2-=1
			continue
		}

		if firstChar < 97 {
			firstChar += 32
		}

		if secondChar < 97 {
			secondChar += 32
		}

		if firstChar != secondChar {
			return false
		}

		pointer1 += 1
		pointer2 -= 1

	}

	return true
}

func isAlphanumeric(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b))
}
