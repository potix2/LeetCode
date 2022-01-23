package explore

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findByVal(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := findByVal(root.Left, val)
	if left != nil {
		return left
	}
	return findByVal(root.Right, val)
}

func lcaForTest(root *TreeNode, p, q int) int {
	node := lowestCommonAncestor(root, findByVal(root, p), findByVal(root, q))
	return node.Val
}

func findPath(root *TreeNode, node *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == node.Val {
		return []*TreeNode{root}
	}
	lhs := findPath(root.Left, node)
	if lhs != nil {
		return append([]*TreeNode{root}, lhs...)
	}
	rhs := findPath(root.Right, node)
	if rhs != nil {
		return append([]*TreeNode{root}, rhs...)
	}
	return nil
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathForP := findPath(root, p)
	pathForQ := findPath(root, q)
	a := pathForP
	b := pathForQ
	if len(a) < len(b) {
		a = pathForQ
		b = pathForP
	}

	lca := a[0]
	for i := 1; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
		lca = a[i]
	}
	return lca
}
