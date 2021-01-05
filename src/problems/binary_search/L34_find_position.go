package binary_search

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := right - (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
