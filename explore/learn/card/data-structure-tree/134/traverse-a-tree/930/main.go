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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	lhs := postorderTraversal(root.Left)
	rhs := postorderTraversal(root.Right)
	return append(append(lhs, rhs...), root.Val)
}
