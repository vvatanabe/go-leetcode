package main

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	var arr []int
	for i, num := range nums {
		if num != i+1 {
			arr = append(arr, i+1)
		}
	}
	return arr
}

func findDisappearedNumbers3(nums []int) []int {
	mp := make(map[int]struct{})
	for _, num := range nums {
		mp[num] = struct{}{}
	}
	var arr []int
	for i := 1; i <= len(nums); i++ {
		_, ok := mp[i]
		if !ok {
			arr = append(arr, i)
		}
	}
	return arr
}

func findDisappearedNumbers2(nums []int) []int {
	var i int
	for i < len(nums) {
		if nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	var arr []int
	for i, _ := range nums {
		if nums[i] != i+1 {
			arr = append(arr, i+1)
		}
	}
	return arr
}
