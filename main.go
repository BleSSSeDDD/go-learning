package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) []string {
	sr := strings.Split(s, "")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		sr[i], sr[j] = sr[j], sr[i]
	}
	return sr
}

func main() {
	a := "qweqwe"
	fmt.Println(a, reverseString(a))
}
