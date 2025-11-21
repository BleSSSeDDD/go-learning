package main

import (
	"fmt"
)

func countPalindromicSubsequence(s string) int {
	var res int
	first := make([]int, 26)
	last := make([]int, 26)

	for i := range first {
		first[i], last[i] = -1, -1
	}

	for i := 0; i < len(s); i++ {
		if first[s[i]-'a'] == -1 {
			first[s[i]-'a'] = i
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if last[s[i]-'a'] == -1 {
			last[s[i]-'a'] = i
		}
	}

	for i := 0; i < len(first); i++ {
		uniq := make(map[byte]struct{})
		if first[i] == -1 || last[i] == -1 {
			continue
		}
		for j := first[i] + 1; j < last[i]; j++ {
			if _, inUse := uniq[s[j]]; !inUse {
				uniq[s[j]] = struct{}{}
				res++
			}
		}
	}

	return res
}

func main() {
	a := "qwqw"
	res := countPalindromicSubsequence(a)
	fmt.Println(res)
}
