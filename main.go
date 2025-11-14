package main

import (
	"fmt"
	"sort"
)

func countPairs(nums []int, target int) int {
	var res int
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] >= target {
			j--
		} else {
			res += j - i
			fmt.Println(i)
			i++
		}
	}
	return res
}

func main() {
	a := []int{1, 1, 2, 2}
	fmt.Println(countPairs(a, 4))
}

(0 1 2 3 4 5 6 7 8)