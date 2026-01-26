package main

import (
	"cmp"
	"fmt"
)

type Heap[T cmp.Ordered] struct {
	buff  []T
	isMin bool
}

func NewMinHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{buff: make([]T, 0), isMin: true}
}

func NewMaxHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{buff: make([]T, 0), isMin: false}
}

// по сути дает ответ на вопрос с кем нам меняться: в случае minHeap возвращает большего, в случае maxHeap меньшего
func compare[T cmp.Ordered](a, b T, isMin bool) bool {
	if isMin {
		return a > b
	}
	return b > a
}

func size[T cmp.Ordered](h *Heap[T]) int {
	if h == nil {
		return 0
	}
	return len(h.buff)
}

func appendToHeap[T cmp.Ordered](h *Heap[T], element T) {
	h.buff = append(h.buff, element)
	hippifyUp(h)
}

func hippifyUp[T cmp.Ordered](h *Heap[T]) {
	index := len(h.buff) - 1
	for index > 0 && compare(h.buff[(index-1)/2], h.buff[index], h.isMin) {
		h.buff[(index-1)/2], h.buff[index] = h.buff[index], h.buff[(index-1)/2]
		index = (index - 1) / 2
	}
}

func deleteFromHeap[T cmp.Ordered](h *Heap[T]) {
	if h == nil || len(h.buff) == 0 {
		return
	}
	h.buff[0], h.buff[len(h.buff)-1] = h.buff[len(h.buff)-1], h.buff[0]
	h.buff = h.buff[0 : len(h.buff)-1]
	hippifyDown(h)
}

func hippifyDown[T cmp.Ordered](h *Heap[T]) {
	index := 0

	for {
		swapIndex := index
		left, right := index*2+1, index*2+2
		if left < len(h.buff) && compare(h.buff[index], h.buff[left], h.isMin) {
			swapIndex = left
		}
		if right < len(h.buff) && compare(h.buff[swapIndex], h.buff[right], h.isMin) {
			swapIndex = right
		}

		if swapIndex == index {
			break
		}

		h.buff[index], h.buff[swapIndex] = h.buff[swapIndex], h.buff[index]
		index = swapIndex
	}
}

func peek[T cmp.Ordered](h *Heap[T]) T {
	if h == nil || len(h.buff) == 0 {
		var zero T
		return zero
	}
	return h.buff[0]
}

func pop[T cmp.Ordered](h *Heap[T]) T {
	if h == nil || len(h.buff) == 0 {
		var zero T
		return zero
	}
	res := h.buff[0]
	deleteFromHeap(h)
	return res
}

func main() {
	// --- Тест Min-Heap ---
	fmt.Println("=== Min-Heap Test ===")
	minHeap := NewMinHeap[int]()

	fmt.Println("Добавляем элементы: 5, 3, 8, 1, 6")
	appendToHeap(minHeap, 5)
	appendToHeap(minHeap, 3)
	appendToHeap(minHeap, 8)
	appendToHeap(minHeap, 1)
	appendToHeap(minHeap, 6)

	fmt.Println("Размер кучи:", size(minHeap))
	fmt.Println("Peek:", peek(minHeap))

	fmt.Println("Pop элементы:")
	for size(minHeap) > 0 {
		fmt.Println(pop(minHeap))
	}

	// --- Тест Max-Heap ---
	fmt.Println("\n=== Max-Heap Test ===")
	maxHeap := NewMaxHeap[int]()

	fmt.Println("Добавляем элементы: 5, 3, 8, 1, 6")
	appendToHeap(maxHeap, 5)
	appendToHeap(maxHeap, 3)
	appendToHeap(maxHeap, 8)
	appendToHeap(maxHeap, 1)
	appendToHeap(maxHeap, 6)

	fmt.Println("Размер кучи:", size(maxHeap))
	fmt.Println("Peek:", peek(maxHeap))

	fmt.Println("Pop элементы:")
	for size(maxHeap) > 0 {
		fmt.Println(pop(maxHeap))
	}

	fmt.Println("\n=== Max-Heap Stress Test (случайные данные) ===")
	stressHeap := NewMaxHeap[int]()
	testData := []int{100, -5, 42, 777, 0, 23, -100, 500, 1, 1, 1, 999}

	fmt.Println("Вставляем в порядке:", testData)
	for _, v := range testData {
		appendToHeap(stressHeap, v)
	}

	fmt.Println("Извлекаем всё (должно быть по убыванию):")
	prev := peek(stressHeap)
	for size(stressHeap) > 0 {
		curr := pop(stressHeap)
		fmt.Printf("%d ", curr)
		// Проверка свойства max-heap: каждый следующий извлечённый должен быть <= предыдущего
		if curr > prev {
			fmt.Printf("\n⛔ ОШИБКА: нарушение порядка! %d > %d\n", curr, prev)
		}
		prev = curr
	}
	fmt.Println("\n✅ Стресс-тест пройден, если выше нет ошибок.")
}
