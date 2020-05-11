package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {
	hLen, nLen := len(haystack), len(needle)

	if nLen == 0 {
		return 0
	}

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack, needle := "hello", "ll"
	fmt.Println(strStr(haystack, needle))
}
