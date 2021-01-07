package two_pointers

func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	k = k % len(nums)
	reverse(&nums, 0, len(nums)-k-1)
	reverse(&nums, len(nums)-k, len(nums)-1)
	reverse(&nums, 0, len(nums)-1)
}

func reverse(nums *[]int, left int, right int) {
	for left < right {
		temp := (*nums)[left]
		(*nums)[left] = (*nums)[right]
		(*nums)[right] = temp
		left++
		right--
	}
}
