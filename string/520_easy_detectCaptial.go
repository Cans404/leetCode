package main

import (
	"fmt"
)

func detectCaptial(str string) bool {
	bytes := []byte(str)
	head := bytes[0:1]
	tail := bytes[1:]

	if isSame(head, 'A', 'Z') {
		return isSame(tail, 'A', 'Z') || isSame(tail, 'a', 'z')
	} else {
		return isSame(tail, 'a', 'z')
	}
}

func isSame(b []byte, start, end byte) bool {
	for _, v := range b {
		if start <= v && v <= end {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(detectCaptial("toDo"))
}
