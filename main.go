package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			k++
		} else {
			nums[k] = nums[i]
		}
	}
	return k
}

func main() {
	a := []int{0, 0, 1, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5, 5, 5, 6, 7}
	fmt.Println(removeDuplicates(a))
}
