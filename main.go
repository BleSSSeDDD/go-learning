package main

import "fmt"

func subarraySum(nums []int, k int) int {
	mapp := make(map[int]int)
	var res int

	sum := 0
	mapp[0]++

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, exists := mapp[sum]; !exists {
			mapp[sum]++
		}

		if mapp[sum/2] > 1 {
			res += mapp[sum/2]
		}
		fmt.Println("i=", i, " s=", sum)
	}

	return res
}

func main() {
	k := 3

	a := []int{1, 2, 3}

	fmt.Println(subarraySum(a, k))
}
