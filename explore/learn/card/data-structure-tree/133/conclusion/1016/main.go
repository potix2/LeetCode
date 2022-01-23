package explore

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	leftmost := root
	for leftmost != nil {
		var nodes []*Node
		head := leftmost
		for head != nil {
			if head.Left != nil {
				nodes = append(nodes, head.Left)
			}
			if head.Right != nil {
				nodes = append(nodes, head.Right)
			}
			head = head.Next
		}

		for i := 1; i < len(nodes); i++ {
			nodes[i-1].Next = nodes[i]
		}

		if len(nodes) > 0 {
			leftmost = nodes[0]
		} else {
			leftmost = nil
		}
	}
	return root
}
