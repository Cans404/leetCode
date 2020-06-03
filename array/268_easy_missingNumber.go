package main

import (
	"fmt"
)

func missingNumber(x []int) int {
	sum := 0
	n := len(x)

	for _, v := range x {
		sum += v
	}

	return n*(n+1)/2 - sum
}

func main() {
	x := []int{9,6,4,2,3,5,7,0,1}

	fmt.Println(missingNumber(x))
}
