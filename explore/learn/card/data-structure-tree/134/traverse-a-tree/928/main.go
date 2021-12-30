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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	lhs := preorderTraversal(root.Left)
	rhs := preorderTraversal(root.Right)
	return append(append([]int{root.Val}, lhs...), rhs...)
}