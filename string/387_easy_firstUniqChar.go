package main

import (
	"fmt"
)

func firstUniqChar(str string) int {
	rec := make([]int, 26)

	for _, v := range str {
		rec[v-'a']++
	}

	for i, v := range str {
		if rec[v-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveflyting"))
}
