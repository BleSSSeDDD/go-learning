package main

import "fmt"

func triangularSum(nums []int) int {
	for len(nums) > 1 {
		for i := 0; i < len(nums)-1; i++ {
			fmt.Println(nums)
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = nums[:len(nums)-1]
	}
	return nums[0]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(triangularSum(a))
}
