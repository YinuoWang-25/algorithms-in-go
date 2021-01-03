package utils

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
