package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countChildren(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return map[bool]int{true: 1}[node.Left != nil] + map[bool]int{true: 1}[node.Right != nil]
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}

	fmt.Println(*root, *root.Left, *root.Right)

	fmt.Println("Структура дерева:")
	fmt.Printf("Корень: %d\n", root.Val)
	fmt.Printf("Левый ребенок: %d\n", root.Left.Val)
	fmt.Printf("Правый ребенок: %d\n", root.Right.Val)

	fmt.Printf("\nУзел %d - лист? %t (нет детей)\n", root.Left.Val, root.Left.Left == nil && root.Left.Right == nil)
	fmt.Printf("Узел %d - лист? %t (нет детей)\n", root.Right.Val, root.Right.Left == nil && root.Right.Right == nil)

	fmt.Println(countChildren(root))
	fmt.Println(countChildren(root.Left))
	fmt.Println(countChildren(root.Right))

}
