package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{9, 2, 34, 54, 53, 0}
	sort.Ints(arr)
	fmt.Println(arr)
}
