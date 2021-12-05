package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func makeBinaryTree(ints []int) *TreeNode {
	root := &TreeNode{
		Val: ints[0],
		Left: makeBinaryTree(ints[1:]),
		Right: makeBinaryTree(ints[2:]),
	}
	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
	return ""
}