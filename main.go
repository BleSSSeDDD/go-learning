package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxD := max(maxDepth(root.Left), maxDepth(root.Right))
	return maxD + 1
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 4}}}}
	fmt.Println(maxDepth(root))
}
