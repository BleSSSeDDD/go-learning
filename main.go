package main

import "cmp"

func main() {

}

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

func compare[T cmp.Ordered](a, b T) bool {
	return a > b
}

func size[T cmp.Ordered](h *Heap[T]) int {
	return len(h.buff)
}

func appendToHeap[T cmp.Ordered](h *Heap[T], element T) {
	h.buff = append(h.buff, element)
	hippifyUp(h.buff, h.isMin)
}

func hippifyUp[T cmp.Ordered](buff []T, isMin bool) {
	index := len(buff) - 1
	if isMin {
		for index > 0 && compare(buff[(index-1)/2], buff[index]) {
			buff[(index-1)/2], buff[index] = buff[index], buff[(index-1)/2]
			index = (index - 1) / 2
		}
	} else {
		for index > 0 && !compare(buff[(index-1)/2], buff[index]) {
			buff[(index-1)/2], buff[index] = buff[index], buff[(index-1)/2]
			index = (index - 1) / 2
		}
	}
}

func deleteFromHeap[T cmp.Ordered](h *Heap[T]) {
	h.buff[0], h.buff[len(h.buff)-1] = h.buff[len(h.buff)-1], h.buff[0]
	h.buff = h.buff[0 : len(h.buff)-1]
	hippifyDown(h.buff, h.isMin)
}

func hippifyDown[T cmp.Ordered](buff []T, isMin bool) {
	index := 0

	if isMin {
		for {
			left, right := index*2+1, index*2+2
			if left < len(buff) && right < len(buff) {
				if compare(buff[left], buff[right]) {
					if compare(buff[index], buff[left]) {
						buff[left], buff[index] = buff[index], buff[left]
						index = left
					} else {
						break
					}
				} else {
					if compare(buff[index], buff[right]) {
						buff[right], buff[index] = buff[index], buff[right]
						index = right
					} else {
						break
					}
				}
			} else if left < len(buff) {
				if compare(buff[index], buff[left]) {
					buff[left], buff[index] = buff[index], buff[left]
					index = left
				} else {
					break
				}
			} else {
				break
			}
		}
	} else {
		for {
			left, right := index*2+1, index*2+2
			if left < len(buff) && right < len(buff) {
				if !compare(buff[left], buff[right]) {
					if !compare(buff[index], buff[left]) {
						buff[left], buff[index] = buff[index], buff[left]
						index = left
					} else {
						break
					}
				} else {
					if !compare(buff[index], buff[right]) {
						buff[right], buff[index] = buff[index], buff[right]
						index = right
					} else {
						break
					}
				}
			} else if left < len(buff) {
				if !compare(buff[index], buff[left]) {
					buff[left], buff[index] = buff[index], buff[left]
					index = left
				} else {
					break
				}
			} else {
				break
			}
		}
	}
}
