package hashtable

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if last, ok := hash[nums[i]]; ok {
			if i-last <= k {
				return true
			}
		}
		hash[nums[i]] = i
	}
	return false
}
