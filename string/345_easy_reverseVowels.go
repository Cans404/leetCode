package main

import (
	"fmt"
)

func isVowel(x byte) bool {
	return x == 'a' || x == 'o' || x == 'e' || x == 'i' || x == 'u' || x == 'A' || x == 'O' || x == 'E' || x == 'I' || x == 'U'
}

func reverseVowels(str string) string {
	length := len(str)
	bytes := []byte(str)
	i, j := 0, length-1

	for i <= j {
		for i <= length-1 && !isVowel(bytes[i]) {
			i++
		}

		for 0 <= j && !isVowel(bytes[j]) {
			j--
		}

		if i <= j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
	}

	return string(bytes)
}

func main() {
	fmt.Println(reverseVowels("LeetCode"))
}
