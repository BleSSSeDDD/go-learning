package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	s := stack{}

	for i := 0; i < len(logs); i++ {
		log := strings.Split(logs[i], ":")
		tmp, _ := strconv.Atoi(log[2])
		id, _ := strconv.Atoi(log[0])

		if log[1] == "start" {
			s.push(tmp)
		} else {
			timeOfFunc := tmp - s.pop()
			res[id] = timeOfFunc
		}
	}

	return res
}

type stack struct {
	buf []int
}

func (s *stack) pop() int {
	tmp := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	fmt.Println("Pop " + strconv.Itoa(tmp))
	return tmp
}

func (s *stack) push(a int) {
	s.buf = append(s.buf, a)
	fmt.Println("Push " + strconv.Itoa(a))
}

func main() {
	a := []string{"0:start:0", "1:start:2", "1:end:5", "2:start:6", "2:end:7", "0:end:8"}
	n := 3
	res := exclusiveTime(n, a)
	fmt.Println(res)
}
