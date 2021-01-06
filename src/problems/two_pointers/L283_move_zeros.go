package two_pointers

func moveZeroes(nums []int) {
	left, cur := 0, 0

	for cur < len(nums) {
		if nums[cur] != 0 {
			nums[left] = nums[cur]
			left = left + 1
		}
		cur = cur + 1
	}
	for left < len(nums) {
		nums[left] = 0
		left = left + 1
	}
}
