package main

import "fmt"

func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle); i++ {
		for j := 0; haystack[i+j] == needle[j]; j++ {
			if j == len(needle)-1 {
				fmt.Println(j)
				return i
			}
		}
	}

	return -1
}

func main() {
	a := strStr("mississippi", "issipi")
	fmt.Println(a)
}
