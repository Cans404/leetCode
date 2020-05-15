package main

import (
	"fmt"
	"math"
)

func thirdMax(x []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, n := range x {
		if n > max1 {
			max1, max2, max3 = n, max1, max2
		}
		switch true {
		case n > max1:
			max1, max2, max3 = n, max1, max2
		case max2 < n && n < max1:
			max2, max3 = n, max2
		case max3 < n && n < max2:
			max3 = n
		}
	}

	if max3 == math.MinInt64 {
		return max1
	}

	return max3
}

func main() {
	nums := []int{2, 2, 3, 1}

	fmt.Println(thirdMax(nums))
}
