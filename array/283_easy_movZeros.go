package main

import (
	"fmt"
)

func moveZeros(x []int) {
	length := len(x)
	counter := 0

	for i := 0; i < length; i++ {
		if x[i] != 0 {
			x[counter] = x[i]
			counter++
		}
	}

	for counter < length {
		x[counter] = 0
		counter++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeros(nums)
	fmt.Printf("%v\n", nums)
}
