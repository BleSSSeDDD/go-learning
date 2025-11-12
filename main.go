package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		fmt.Println(l, r)
		if nums[l+(r-l)/2] < target {
			l = l + (r-l)/2
		} else if nums[l+(r-l)/2] > target {
			r = l + (r-l)/2
		} else {
			break
		}
	}
	return l
}

func main() {
	a := []int{1, 2, 3, 3, 3, 4, 24356}
	fmt.Println(searchInsert(a, 4))
}
