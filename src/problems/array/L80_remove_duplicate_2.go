package array

func removeDuplicatesII(nums []int) int {
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow = slow + 1
		}
		fast = fast + 1
	}
	return slow
}
