package explore

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	preorder []int
	inorder  []int
	indexMap map[int]int
}

func (s *Solution) helper(inLeft, inRight int) *TreeNode {
	if inLeft > inRight {
		return nil
	}
	v := s.preorder[0]
	s.preorder = s.preorder[1:]
	idx := s.indexMap[v]
	lhs := s.helper(inLeft, idx-1)
	rhs := s.helper(idx+1, inRight)
	return &TreeNode{
		Val:   v,
		Left:  lhs,
		Right: rhs,
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	imap := map[int]int{}
	for i, v := range inorder {
		imap[v] = i
	}
	s := &Solution{
		preorder: preorder,
		inorder:  inorder,
		indexMap: imap,
	}
	return s.helper(0, len(inorder)-1)
}
