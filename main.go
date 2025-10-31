package main

import (
	"fmt"
	"sort"
)

func TwoSum(nums []int, target int) []int {

	mapp := make(map[int]int)

	result := make([]int, 2)
	var exists bool

	for i, num := range nums {
		pair := target - num
		result[0] = i
		result[1], exists = mapp[pair]
		mapp[num] = i
		if exists {
			return result
		}
	}

	return nil
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	mapp := make(map[[3]int]struct{})

	for i, num := range nums {
		tmp1 := make([]int, 1)

		tmp1[0] = i

		tmp2 := TwoSum(nums[i+1:], -num)
		if tmp2 != nil {
			tmp2[0] += i + 1
			tmp2[1] += i + 1
		}

		if tmp2 != nil {
			tmp1 = append(tmp1, tmp2...)
			sort.Ints(tmp1)
			triple := [3]int{
				nums[i],
				nums[tmp2[0]],
				nums[tmp2[1]],
			}

			if _, exists := mapp[triple]; !exists {
				res = append(res, tmp1)
				mapp[triple] = struct{}{}
			} else {
				fmt.Println("Уже есть ", triple)
			}
		}
	}
	return res
}

func main() {
	nums := []int{-4, -2, -2, -1, -1, 0, 1, 2, 2, 3, 4, 5}
	fmt.Println(threeSum(nums))
}
