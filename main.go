package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 15)
	sort.Ints(nums)
	var s int
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			s = nums[l] + nums[r]
			if s == -nums[i] {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l-1] == nums[l] {
					l++
				}
				for r > l && nums[r+1] == nums[r] {
					r--
				}
			} else if s > -nums[i] {
				r--
				for r > l && nums[r+1] == nums[r] {
					r--
				}
			} else {
				l++
				for l < r && nums[l-1] == nums[l] {
					l++
				}
			}

		}
	}
	return res
}

func main() {
	a := []int{-100, 50, 50, 50, 50, 70, 30}
	fmt.Println(threeSum(a))
}
