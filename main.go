package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {

	var i int

	for i < len(bits) {
		if bits[i] == 0 {
			if i == len(bits)-1 {
				return true
			}
			i++
		} else {
			i += 2
			if i >= len(bits) {
				return false
			}
		}
	}
	return true
}

func main() {
	a := []int{0, 1, 0, 1, 0}

	fmt.Println(isOneBitCharacter(a))
}
