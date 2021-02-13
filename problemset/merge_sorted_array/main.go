package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var cur int = len(nums1)
	for m > 0 && n > 0 {
		cur--
		if nums1[m-1] >= nums2[n-1] {
			m--
			nums1[cur] = nums1[m]
		} else {
			n--
			nums1[cur] = nums2[n]
		}
	}
	for n > 0 {
		n--
		cur--
		nums1[cur] = nums2[n]
	}
}
