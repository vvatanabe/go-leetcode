package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	max := len(nums)
	if max < 3 {
		return [][]int{}
	}
	var output [][]int
	sort.Ints(nums)
	uniq := make(map[int]struct{})
	for i := 0; i < max; i++ {
		uniq[nums[i]] = struct{}{}
	}
	if len(uniq) == 1 {
		if _, ok := uniq[0]; ok && max > 2 {
			return [][]int{{0, 0, 0}}
		}
	}
	for i := 0; i < max; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := max - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k > max-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			total := nums[i] + nums[j] + nums[k]
			if total > 0 {
				k--
			} else if total < 0 {
				j++
			} else {
				output = append(output, []int{nums[i], nums[j], nums[k]})
				k--
				j++
			}
		}
	}
	return output
}

func threeSum2(nums []int) [][]int {
	max := len(nums)
	if max < 3 {
		return [][]int{}
	}
	var output [][]int
	sort.Ints(nums)
	for i := 0; i < max-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < max-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < max; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					match := []int{nums[i], nums[j], nums[k]}
					output = append(output, match)
					break
				}
			}
		}
	}
	return output
}
