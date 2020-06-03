package main

import (
	"fmt"
)

func majorityElement(x []int) int {
	res, cnt := 0, 0

	for _, v := range x {
		if cnt == 0 {
			res = v
			cnt++
		} else if res == v {
			cnt++
		} else {
			cnt--
		}
	}

	return res
}

func main() {
	x := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(x))
}
