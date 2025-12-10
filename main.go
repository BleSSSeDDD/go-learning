package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{{20, 30}, {25, 10, 0}, {-10, 20}, {-10, 20}, {-10, 20}}
	fmt.Println(a)

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	fmt.Println(a)
}
