package explore

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalSubtree(root *TreeNode, parentVal int) bool {
	if root == nil {
		return true
	}
	return root.Val == parentVal &&
		isUnivalSubtree(root.Left, parentVal) &&
		isUnivalSubtree(root.Right, parentVal)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if isUnivalSubtree(root, root.Val) {
		return 1 + countUnivalSubtrees(root.Left) + countUnivalSubtrees(root.Right)
	} else {
		return countUnivalSubtrees(root.Left) + countUnivalSubtrees(root.Right)
	}
}
