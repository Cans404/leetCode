package main

import (
	"fmt"
)

func searchInsert(x []int, t int) int {
	length := len(x)

	if length == 0 {
		return 0
	}

	for i, v := range x {
		if v >= t {
			return i
		}
	}

	return length
}

func main() {
	nums, target := []int{1, 3, 5, 6}, 5

	fmt.Println(searchInsert(nums, target))
}
