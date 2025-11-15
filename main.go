package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	k := 0
	i := 0
	for i < len(nums)-k {
		if nums[i] == 0 {
			for j := i; j < len(nums)-1-k; j++ {
				nums[j] = nums[j+1]
				fmt.Println(nums)
			}
			nums[len(nums)-1-k] = 0
			k++
		} else {
			i++
		}
	}
}

func main() {
	a := []int{0, 0, 1, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5, 5, 5, 6, 7}
	moveZeroes(a)
	fmt.Println(a)
}
