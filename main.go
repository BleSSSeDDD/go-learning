package main

import "fmt"

func main() {
	a := &TreeNode{Val: "a"}
	b := &TreeNode{Val: "b"}
	c := &TreeNode{Val: "c"}
	d := &TreeNode{Val: "d"}
	e := &TreeNode{Val: "e"}
	f := &TreeNode{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	fmt.Println(a.DFS())
}

func (t *TreeNode) DFS() []string {
	if t == nil {
		return nil
	}

	res := make([]string, 0)
	s := Stack[*TreeNode]{}

	s.Push(t)

	for !s.Empty() {
		curr := s.Pop()
		res = append(res, curr.Val)

		if curr.Right != nil {
			s.Push(curr.Right)
		}

		if curr.Left != nil {
			s.Push(curr.Left)
		}
	}

	return res
}

type TreeNode struct {
	Val   string
	Right *TreeNode
	Left  *TreeNode
}

func (s *Stack[T]) Empty() bool {
	return len(s.buff) == 0
}

type Stack[T any] struct {
	buff []T
}

func (s *Stack[T]) Pop() T {
	if len(s.buff) == 0 {
		var zero T
		return zero
	}
	res := s.buff[len(s.buff)-1]
	s.buff = s.buff[0 : len(s.buff)-1]
	return res
}

func (s *Stack[T]) Push(a T) {
	s.buff = append(s.buff, a)
}
