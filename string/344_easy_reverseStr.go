package main

import (
	"fmt"
)

func main() {
	str := "hello, world"
	bytes := []byte(str)

	i, j := 0, len(bytes) - 1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	fmt.Println(string(bytes))
}
