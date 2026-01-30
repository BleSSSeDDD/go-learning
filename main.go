package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	rotateRight(a, 2)
	fmt.Println(a)

	a = []int{0, 1, 2, 3, 4}
	rotateLeft(a, 2)
	fmt.Println(a)
}

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

// для ротейта влево мы переворачиваем первые k, потом остальные, потом переворачиваем все целиком
func rotateLeft(a []int, k int) {
	k = k % len(a)
	reverse(a[0:k])
	reverse(a[k:])
	reverse(a)

}

// для ротейта вправо мы переворачиваем всё целиком, потом первые k, потом остальные
func rotateRight(a []int, k int) {
	k = k % len(a)
	reverse(a)
	reverse(a[k:])
	reverse(a[0:k])
}
