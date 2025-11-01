package main

import "fmt"

func triangularSum(nums []int) int {
	tmpSlc := make([][]int, 1, 10)
	tmpSlc[0] = nums
	for i := 1; i < len(nums); i++ {
		tmp := make([]int, len(nums)-i+1, len(nums)-i+1)
		for j := 0; j < len(nums)-i; j++ {
			fmt.Println(i, j, tmpSlc)
			tmp[j] = tmpSlc[i-1][j] + tmpSlc[i-1][j+1]
		}
		tmpSlc = append(tmpSlc, tmp)
	}
	return tmpSlc[len(tmpSlc)-1][0]
}

func main() {
	fmt.Println(triangularSum([]int{1, 2, 3, 4, 5}))
}
