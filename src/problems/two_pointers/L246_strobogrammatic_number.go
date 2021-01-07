package two_pointers

func isStrobogrammatic(num string) bool {
	hash := map[byte]byte{
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
		'0': '0',
	}

	bytes := []byte(num)
	left, right := 0, len(bytes)-1
	for left <= right {
		if hash[num[left]] != num[right] {
			return false
		}
		left++
		right--
	}
	return true
}
