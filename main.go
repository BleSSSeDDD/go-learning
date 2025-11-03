package main

import (
	"fmt"
)

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
func majorityElement(nums []int) int {
	mapp := make(map[int]int)
	maxElem, maxCount := nums[0], 0
	mapp[nums[0]]++
	for i := 1; i < len(nums); i++ {
		mapp[nums[i]]++
		if mapp[nums[i]] > maxCount {
			maxCount = mapp[nums[i]]
			maxElem = nums[i]
		}
	}
	return maxElem
}

func main() {
	a := []int{3, 3, 4}
	fmt.Println(majorityElement(a))
}
