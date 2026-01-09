package main

import (
	"fmt"
	"sort"
)

func binsearch(a []int, target int, l int, r int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2

	if a[mid] == target {
		return mid
	} else if a[mid] < target {
		return binsearch(a, target, mid+1, r)
	}

	return binsearch(a, target, l, mid-1)
}

func main() {
	a := []int{}
	sort.Ints(a)
	fmt.Println(binsearch(a, 228, 0, len(a)-1))
}
