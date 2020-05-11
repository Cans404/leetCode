package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	tag := shortest(strs)

	if tag == "" {
		return ""
	}

	for i := 0; i < len(tag); i++ {
		s := tag[0 : len(tag)-i]
		for _, str := range strs {
			if s != str[0:len(s)] {
				s = ""
				break
			}
		}

		if s != "" {
			return s
		}
	}
	
	return ""
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s := strs[0]
	for _, str := range strs {
		if len(str) < len(s) {
			s = str
		}
	}
	return s
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Printf("\"%s\"\n", longestCommonPrefix(strs))
}
