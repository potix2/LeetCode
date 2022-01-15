package explore

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	postorder []int
	inorder   []int
	indexMap  map[int]int
}

func (s *Solution) helper(inLeft, inRight int) *TreeNode {
	if inLeft > inRight {
		return nil
	}

	v := s.postorder[len(s.postorder)-1]
	s.postorder = s.postorder[:len(s.postorder)-1]
	idx := s.indexMap[v]
	rhs := s.helper(idx+1, inRight)
	lhs := s.helper(inLeft, idx-1)
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	idx := map[int]int{}
	for i, v := range inorder {
		idx[v] = i
	}
	solution := &Solution{
		inorder:   inorder,
		postorder: postorder,
		indexMap:  idx,
	}
	return solution.helper(0, len(inorder)-1)
}
