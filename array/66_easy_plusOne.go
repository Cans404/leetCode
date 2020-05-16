package main

import (
	"fmt"
)

func plusOne(x []int) []int {
	length := len(x)

	x[length-1]++

	for i := length - 1; i > 0; i-- {
		if x[i] <= 9 {
			break
		} else {
			x[i] -= 10
			x[i-1] ++
		}
	}

	if x[0] > 9 {
		x[0] -= 10
		x = append([]int{1}, x...)
	}

	return x
}

func main() {
	arr := []int{9, 9, 9, 9}
	fmt.Printf("%v\n", plusOne(arr))
}
