package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = struct{}{}
	}
	var out []int
	for i := 0; i < len(nums2); i++ {
		_, ok := m[nums2[i]]
		if ok {
			out = append(out, nums2[i])
			delete(m, nums2[i])
		}
	}
	return out
}
