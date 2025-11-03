package main

import (
	"fmt"
	"strings"
)

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string
// by deleting some (can be none) of the characters without disturbing the relative positions of the remaining
// characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	cnt := 0
	tres := strings.Split(t, "")
	sres := strings.Split(s, "")
	for i := 0; i < len(tres); i++ {
		if sres[cnt] == tres[i] {
			if cnt == len(sres)-1 {
				return true
			}
			cnt++
		}
	}
	return false
}

func main() {
	a := ""
	b := "a"
	fmt.Println(isSubsequence(a, b))
}
