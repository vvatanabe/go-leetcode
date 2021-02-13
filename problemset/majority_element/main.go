package main

func majorityElement(nums []int) int {
	var major, num int
	for _, v := range nums {
		switch {
		case v != major && num <= 0:
			major = v
		case v != major:
			num--
		default:
			num++
		}
	}
	return major
}

func majorityElement2(nums []int) int {
	h := len(nums) / 2
	m := make(map[int]int)
	var major, max int
	for _, v := range nums {
		cnt := m[v]
		cnt++
		m[v] = cnt
		if cnt > h {
			major = v
			break
		}
		if max < cnt {
			max = cnt
			major = v
		}
	}
	return major
}
