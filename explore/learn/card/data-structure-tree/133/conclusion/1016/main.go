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

/*
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
*/
func processChild(childNode, prev, leftmost *Node) (*Node, *Node) {
	if childNode != nil {
		if prev != nil {
			prev.Next = childNode
		} else {
			leftmost = childNode
		}
		prev = childNode
	}
	return prev, leftmost
}

func connect(root *Node) *Node {
	if root != nil {
		return nil
	}
	leftmost := root
	for leftmost != nil {
		var prev *Node
		cur := leftmost
		leftmost = nil
		for cur != nil {
			prev, leftmost = processChild(cur.Left, prev, leftmost)
			prev, leftmost = processChild(cur.Right, prev, leftmost)
			cur = cur.Next
		}
	}
	return root
}
