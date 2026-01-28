package main

import (
	"fmt"

	"github.com/BleSSSeDDD/go-learning/trees"
)

func main() {

	a := &trees.TreeNode[string]{Val: "a"}
	b := &trees.TreeNode[string]{Val: "b"}
	c := &trees.TreeNode[string]{Val: "c"}
	d := &trees.TreeNode[string]{Val: "d"}
	e := &trees.TreeNode[string]{Val: "e"}
	f := &trees.TreeNode[string]{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	fmt.Println(a.IterativeDFS())
	fmt.Println(a.RecursiveDFS())
	fmt.Println(a.BFS())

	fmt.Println(a.IncludesBFS("f"))
	fmt.Println(a.IncludesBFS("g"))

	fmt.Println(a.IncludesIterativeDFS("f"))
	fmt.Println(a.IncludesIterativeDFS("g"))

	fmt.Println(a.IncludesRecursiveDFS("f"))
	fmt.Println(a.IncludesRecursiveDFS("g"))

}
