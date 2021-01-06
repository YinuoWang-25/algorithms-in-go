package binary_search

func missingElement(nums []int, k int) int {
	left, right := 0, len(nums)-1
	missing := nums[right] - nums[left] - (right - left)

	if k > missing {
		return nums[right] + k - missing
	}

	for left < right+1 {
		mid := left + (right-left)/2
		missing := nums[mid] - nums[left] - (mid - left)
		if k > missing {
			k -= missing
			left = mid
		} else {
			right = mid
		}
	}
	return nums[left] + k
}
