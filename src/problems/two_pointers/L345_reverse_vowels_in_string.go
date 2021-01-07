package two_pointers

func reverseVowels(s string) string {
	bytes := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isVowels(bytes[left]) {
			left++
		}
		for left < right && !isVowels(bytes[right]) {
			right--
		}
		temp := bytes[left]
		bytes[left] = bytes[right]
		bytes[right] = temp
		left++
		right--

	}
	return string(bytes)
}

func isVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
