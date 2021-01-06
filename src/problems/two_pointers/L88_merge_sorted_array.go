package two_pointers

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[p] = nums1[i]
			i = i - 1
		} else {
			nums1[p] = nums2[j]
			j = j - 1
		}
		p = p - 1
	}
	for j >= 0 {
		nums1[p] = nums2[j]
		j = j - 1
		p = p - 1
	}
}
