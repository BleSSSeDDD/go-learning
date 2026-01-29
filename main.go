package main

import "fmt"

func main() {
	a := []int{4, 1, -1, 2, -1, 2, 3}

	fmt.Println(topKFrequent(a, 2))
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, k)

	freq := make(map[int]int)

	l := 0

	for r := range nums {
		if _, exists := freq[nums[r]]; !exists {
			freq[nums[r]]++
			nums[r], nums[l] = nums[l], nums[r]
			l++
		} else {
			freq[nums[r]]++
		}
	}

	numsWithFreq := make([][]int, l)

	h := &heap{}

	for i := range l {

		numsWithFreq[i] = make([]int, 2)

		numsWithFreq[i][0] = nums[i]
		numsWithFreq[i][1] = freq[nums[i]]

		if len(h.buff) == k {
			curr := h.pop()

			if curr[1] >= numsWithFreq[i][1] {
				h.push(curr)
			} else {
				h.push(numsWithFreq[i])
			}
		} else {
			h.push(numsWithFreq[i])
		}

		fmt.Println(numsWithFreq)

		fmt.Println("heap: ", h.buff)

	}

	for i := range res {
		pair := h.pop()
		if pair == nil {
			res[i] = -1
			continue
		}
		res[i] = pair[0]
	}

	return res
}

type heap struct {
	buff [][]int
}

func (h *heap) push(arr []int) {
	h.buff = append(h.buff, arr)
	hippifyUp(h)
}

func (h *heap) pop() []int {
	if len(h.buff) == 0 {
		return nil
	}

	res := h.buff[0]

	h.buff[0], h.buff[len(h.buff)-1] = h.buff[len(h.buff)-1], h.buff[0]
	h.buff = h.buff[0 : len(h.buff)-1]

	hippifyDown(h)

	return res
}

func hippifyUp(h *heap) {
	i := len(h.buff) - 1

	for i > 0 {
		if h.buff[i][1] < h.buff[(i-1)/2][1] {
			h.buff[i], h.buff[(i-1)/2] = h.buff[(i-1)/2], h.buff[i]
			i = (i - 1) / 2
		} else {
			break
		}
	}
}

func hippifyDown(h *heap) {
	i := 0

	for {
		swapIndex := i

		if i*2+1 < len(h.buff) && h.buff[i][1] > h.buff[i*2+1][1] {
			swapIndex = i*2 + 1
		}
		if i*2+2 < len(h.buff) && h.buff[swapIndex][1] > h.buff[i*2+2][1] {
			swapIndex = i*2 + 2
		}

		if swapIndex == i {
			break
		}

		h.buff[i], h.buff[swapIndex] = h.buff[swapIndex], h.buff[i]

		i = swapIndex
	}
}
