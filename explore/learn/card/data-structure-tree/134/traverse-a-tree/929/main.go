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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	lhs := inorderTraversal(root.Left)
	rhs := inorderTraversal(root.Right)
	return append(append(lhs, root.Val), rhs...)
}
