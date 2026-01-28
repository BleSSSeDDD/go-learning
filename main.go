package main

import (
	"fmt"
)

func main() {
	a := &TreeNode[string]{Val: "a"}
	b := &TreeNode[string]{Val: "b"}
	c := &TreeNode[string]{Val: "c"}
	d := &TreeNode[string]{Val: "d"}
	e := &TreeNode[string]{Val: "e"}
	f := &TreeNode[string]{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	fmt.Println(a.IterativeDFS())
	fmt.Println(a.RecursiveDFS())
	fmt.Println(a.BFS())

}

func (t *TreeNode[T]) BFS() []T {
	res := make([]T, 0)

	q := Queue[*TreeNode[T]]{}
	q.Push(t)

	for !q.Empty() {
		curr := q.Pop()

		res = append(res, curr.Val)

		if curr.Left != nil {
			q.Push(curr.Left)
		}

		if curr.Right != nil {
			q.Push(curr.Right)
		}
	}

	return res
}

func (t *TreeNode[T]) RecursiveDFS() []T {
	res := make([]T, 0)

	if t == nil {
		return []T{}
	}

	res = append(res, t.Val)
	res = append(res, t.Left.RecursiveDFS()...)
	res = append(res, t.Right.RecursiveDFS()...)

	return res
}

func (t *TreeNode[T]) IterativeDFS() []T {
	if t == nil {
		return nil
	}

	res := make([]T, 0)
	s := Stack[*TreeNode[T]]{}

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

type TreeNode[T any] struct {
	Val   T
	Right *TreeNode[T]
	Left  *TreeNode[T]
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

type Queue[T any] struct {
	buff []T
}

func (q *Queue[T]) Push(a T) {
	q.buff = append(q.buff, a)
}

func (q *Queue[T]) Pop() T {
	var zero T
	if q.Empty() {
		return zero
	}
	res := q.buff[0]
	q.buff = q.buff[1:]
	return res
}

func (q *Queue[T]) Empty() bool {
	return len(q.buff) == 0
}
