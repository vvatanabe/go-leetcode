package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		mp[nums2[i]] = i
	}
	arr := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		next := -1
		if v, ok := mp[nums1[i]]; ok {
			for j := v; j < len(nums2); j++ {
				if nums1[i] < nums2[j] {
					next = nums2[j]
					break
				}
			}
		}
		arr[i] = next
	}
	return arr
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	arr := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		next := -1
		for j := 0; j < len(nums2); j++ {
			if nums1[i] != nums2[j] {
				continue
			}
			for l := j + 1; l < len(nums2); l++ {
				if nums1[i] < nums2[l] {
					next = nums2[l]
					break
				}
			}
		}
		arr[i] = next
	}
	return arr
}
