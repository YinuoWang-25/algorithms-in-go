package binary_search

func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		val := guess(mid)
		if val == 0 {
			return mid
		} else if val == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func guess(num int) int {
	return -1
}
