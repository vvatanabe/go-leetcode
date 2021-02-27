package main

import "fmt"

// root:
//   i = 1
//
// parent:
//   parent(i) = i / 2
//
// left child node:
//   left(i) = 2i
//
// right child node:
//   right(i) = 2i + 1
//
// leaf nodes:
//   after n/2+1

func main() {
	nums := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("nums:", nums)
	heapSort(nums)
	fmt.Println("nums:", nums)
}

func heapSort(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
	last := len(arr) - 1
	for i := last; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, root, size int) {
	max := root
	l := 2*root + 1
	r := 2*root + 2
	if l < size && arr[max] < arr[l] {
		max = l
	}
	if r < size && arr[max] < arr[r] {
		max = r
	}
	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		heapify(arr, max, size)
	}
}
