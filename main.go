package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	var sum int
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum-target == 0 {
				return target
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {
	a := []int{-100, 50, 50, 50, 50, 70, 30}
	fmt.Println(threeSumClosest(a, 228))
}
