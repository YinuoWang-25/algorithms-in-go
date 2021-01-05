package binary_search

func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := right - (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
