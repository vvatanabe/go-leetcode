package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = m[nums1[i]] + 1
	}
	var out []int
	for i := 0; i < len(nums2); i++ {
		cnt, ok := m[nums2[i]]
		if ok && cnt > 0 {
			out = append(out, nums2[i])
			m[nums2[i]] = m[nums2[i]] - 1
		}
	}
	return out
}
