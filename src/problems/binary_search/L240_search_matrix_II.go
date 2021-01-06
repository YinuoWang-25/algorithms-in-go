package binary_search

func searchMatrixII(matrix [][]int, target int) bool {
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
	row, col := n-1, 0
	for row >= 0 && col < m {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			col = col + 1
		} else {
			row = row - 1
		}
	}
	return false
}
