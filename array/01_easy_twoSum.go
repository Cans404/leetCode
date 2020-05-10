// 给定一个整数数组nums和一个目标值target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
package main

import (
	"fmt"
	"os"
)

func sliceToMap(s []int) map[int]int {
	m := make(map[int]int, len(s))
	for i, v := range s {
		m[v] = i
	}
	return m
}

func main() {
	target := 17 
	s := []int{2, 7, 11, 15}

	m := sliceToMap(s)
	for i, v := range s {
		mv, ok := m[target-v]
		if ok {
			fmt.Printf("%d: %d\n%d: %d\n", i, v, mv, s[mv])
			os.Exit(0)
		}
	}

	fmt.Println("not found.")
}
