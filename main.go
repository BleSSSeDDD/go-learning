package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0, 10)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		fmt.Println("i = ", i)
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			fmt.Println("j = ", j)
			l, r := j+1, len(nums)-1
			for l < r {
				fmt.Println("l = ", l, " r = ", r)
				if nums[l]+nums[r]+nums[i]+nums[j] == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					r--
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else if nums[l]+nums[r] > target-(nums[i]+nums[j]) {
					r--
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else {
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				}
			}
		}
	}
	return res
}

func main() {
	a := []int{-2, -1, 0, 0, 1, 2}

	fmt.Println(fourSum(a, 0))
}
