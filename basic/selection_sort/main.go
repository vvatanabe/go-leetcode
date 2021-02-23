package main

import "fmt"

// 数列をソートするアルゴリズム。
// 数列を線形探索して最小値を探す。
// 最小値を列の左端と交換する。
// 同様の操作を全ての数字がソート済みになるまで繰り返す。

func main() {
	nums := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("nums:", nums)
	selectionSort(nums)
	fmt.Println("nums:", nums)
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
