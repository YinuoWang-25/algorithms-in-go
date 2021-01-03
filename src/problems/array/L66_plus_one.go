package array

func plusOne(digits []int) []int {
	var res []int

	sum, idx := 1, len(digits)-1

	for sum > 0 || idx >= 0 {
		if idx >= 0 {
			sum += digits[idx]
		}
		res = append([]int{sum % 10}, res...)
		idx = idx - 1
		sum = sum / 10
	}
	return res
}
