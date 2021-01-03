package array

func removeDuplicates(nums []int) int {
	slow, fast := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow = slow + 1
		}
		fast = fast + 1
	}
	return slow
}
