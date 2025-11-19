package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	operations := make(map[string]struct{})
	operations["+"] = struct{}{}
	operations["-"] = struct{}{}
	operations["*"] = struct{}{}
	operations["/"] = struct{}{}
	s := stack{}
	for _, str := range tokens {
		if _, isOp := operations[str]; isOp {
			if str == "+" {
				s.push(s.pop() + s.pop())
			}
			if str == "-" {
				a := s.pop()
				b := s.pop()
				s.push(b - a)
			}
			if str == "*" {
				s.push(s.pop() * s.pop())
			}
			if str == "/" {
				a := s.pop()
				b := s.pop()
				s.push(b / a)
			}
		} else {
			num, _ := strconv.Atoi(str)
			s.push(num)
		}
	}

	return s.pop()
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
	a := []string{"4", "13", "5", "/", "+"}
	res := evalRPN(a)
	fmt.Println(res)
}
