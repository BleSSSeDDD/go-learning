package main

import (
	"fmt"
)

func isHappy(n int) bool {
	mapp := make(map[int]struct{})
	var squaredDigits int
	for n != 1 {
		if _, exists := mapp[n]; exists {
			return false
		}
		mapp[n] = struct{}{}
		squaredDigits = 0
		for n > 0 {
			squaredDigits += (n % 10) * (n % 10)
			n /= 10
		}
		n = squaredDigits
	}
	return true
}

func main() {
	fmt.Println(isHappy(2))
}
