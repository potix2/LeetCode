package main

type TreeNode struct {
	Val int
	Left *TreeNode
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lhs := maxDepth(root.Left)
	rhs := maxDepth(root.Right)
	if lhs < rhs {
		return rhs + 1
	} else {
		return lhs + 1
	}
}