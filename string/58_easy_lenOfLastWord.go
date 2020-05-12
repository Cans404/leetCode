package main

import (
	"fmt"
)

func lenOfLastWord(str string) int {
	sLen := len(str)

	if sLen == 0 {
		return 0
	}

	counter := 0
	for i := sLen - 1; i >= 0; i-- {
		if str[i] == ' ' {
			if counter != 0 {
				return counter
			} else {
				continue
			}

		} else {
			counter++
		}
	}

	return 0
}

func main() {
	s := "wow awesome amazing incredible gorgeous exciting"

	fmt.Println(lenOfLastWord(s))
}
