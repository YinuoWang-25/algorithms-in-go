package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	left, right := 0, n*m-1

	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
