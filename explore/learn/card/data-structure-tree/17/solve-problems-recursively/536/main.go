package explore

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricNode(lhs *TreeNode, rhs *TreeNode) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs != nil && rhs != nil {
		return lhs.Val == rhs.Val && isSymmetricNode(lhs.Left, rhs.Right) && isSymmetricNode(lhs.Right, rhs.Left)
	}
	return false
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricNode(root.Left, root.Right)
}