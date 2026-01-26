package main

import (
	"cmp"
	"fmt"
)

func main() {

	fmt.Println("++++++++1 3 2 9 -10 -> -10 1 2 3 9+++++++++++")
	fmt.Println("+++MaxHeap+++")
	h := NewMaxHeap[int]()

	h.Append(1)
	h.Append(3)
	h.Append(2)
	h.Append(9)
	h.Append(-10)

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	fmt.Println("+++MinHeap+++")

	h = NewMinHeap[int]()

	h.Append(1)
	h.Append(3)
	h.Append(2)
	h.Append(9)
	h.Append(-10)

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	fmt.Println("\n=== СТРЕСС-ТЕСТ MaxHeap ===")
	h2 := NewMaxHeap[int]()
	testData := []int{100, -50, 0, 777, 23, -777, 42, 1}

	fmt.Println("Вставляем:", testData)
	for _, v := range testData {
		h2.Append(v)
	}

	fmt.Println("Извлекаем (должно быть по убыванию):")
	prev := h2.Pop()
	fmt.Print(prev, " ")
	for i := 0; i < len(testData)-1; i++ {
		curr := h2.Pop()
		fmt.Print(curr, " ")
		if curr > prev { // В max-heap каждый следующий должен быть меньше предыдущего
			fmt.Printf("\n⛔ Ошибка порядка! %d > %d\n", curr, prev)
		}
		prev = curr
	}
	fmt.Println("\n✅ Тест пройден, если нет ошибок выше")
}

type Heap[T cmp.Ordered] struct {
	data  []T
	isMin bool
}

func NewMinHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{data: make([]T, 0), isMin: true}
}

func NewMaxHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{data: make([]T, 0), isMin: false}
}

func compare[T cmp.Ordered](a, b T, isMin bool) bool {
	if isMin {
		return a > b
	} else {
		return b > a
	}
}

func (h *Heap[T]) Append(newElement T) {
	h.data = append(h.data, newElement)
	hippifyUp(h)
}

func hippifyUp[T cmp.Ordered](h *Heap[T]) {
	index := len(h.data) - 1
	for index > 0 && compare(h.data[(index-1)/2], h.data[index], h.isMin) {
		h.data[index], h.data[(index-1)/2] = h.data[(index-1)/2], h.data[index]
		index = (index - 1) / 2
	}
}

func (h *Heap[T]) Pop() T {
	if len(h.data) == 0 {
		var zero T
		return zero
	}

	res := h.data[0]

	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[0 : len(h.data)-1]

	hippifyDown(h)
	return res
}

func hippifyDown[T cmp.Ordered](h *Heap[T]) {
	index := 0
	for {
		swapIndex := index

		left, right := index*2+1, index*2+2

		if left < len(h.data) && compare(h.data[index], h.data[left], h.isMin) {
			swapIndex = left
		}
		if right < len(h.data) && compare(h.data[swapIndex], h.data[right], h.isMin) {
			swapIndex = right
		}

		if swapIndex == index {
			break
		}

		h.data[index], h.data[swapIndex] = h.data[swapIndex], h.data[index]
		index = swapIndex
	}
}
