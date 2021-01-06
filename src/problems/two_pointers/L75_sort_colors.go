package two_pointers

func sortColors(nums []int)  {
	zero, one, two, cur := 0, 0,len(nums) - 1, 0
	for cur <= two {
		if nums[cur] == 0 {
			swap(&nums, cur, zero)
			zero = zero + 1
			cur = cur + 1
		} else if nums[cur] == 1 {
			one = one + 1
			cur = cur + 1
		} else {
			swap(&nums, cur, two)
			two  = two - 1
		}
	}
}

func swap(nums *[]int, p int, q int) {
	temp := (*nums)[p]
	(*nums)[p] = (*nums)[q]
	(*nums)[q] = temp
}