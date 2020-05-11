package main

import (
	"fmt"
	"os"
)

func main() {
	s := []int {1, 2, 3, 1}
	m := make(map[int]bool)
	
	for _, v := range s {
		mv, ok := m[v]
		if ok {
			fmt.Println(mv)
			os.Exit(0)		
		} else {
			m[v] = true 
		}
	}

	fmt.Println(false)
}
