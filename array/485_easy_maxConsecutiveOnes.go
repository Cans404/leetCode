package main

import (
	"fmt"
)

func maxConsecutiveOnes(x []int) int {
	counter, max := 0, 0

	for _, v := range x {
		if v == 1 {
			counter++
		}

		if counter > max {
			max = counter
		}

		if v == 0 {
			counter = 0
		}
	}

	return max
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}

	fmt.Println(maxConsecutiveOnes(nums))
}
