package array

import "strconv"

func addBinary(a string, b string) string {
	var res string
	i, j, sum := len(a)-1, len(b)-1, 0
	for i >= 0 || j >= 0 || sum > 0 {
		if i >= 0 {
			n1, _ := strconv.ParseInt(string(a[i]), 10, 32)
			sum = sum + (int)(n1)
			i = i - 1
		}
		if j >= 0 {
			n2, _ := strconv.ParseInt(string(b[j]), 10, 32)
			sum = sum + (int)(n2)
			j = j - 1
		}
		numStr := strconv.Itoa(sum % 2)
		sum = sum / 2
		res = numStr + res
	}
	return res
}
