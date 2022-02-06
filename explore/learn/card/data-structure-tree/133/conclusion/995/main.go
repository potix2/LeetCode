package explore

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
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

type Codec struct {
	levels []string
	nodes  []*TreeNode
	pos    int
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) helper(node *TreeNode) {
	this.levels = append(this.levels, fmt.Sprintf("%d", node.Val))
	if node.Left != nil {
		this.helper(node.Left)
	} else {
		this.levels = append(this.levels, "N")
	}
	if node.Right != nil {
		this.helper(node.Right)
	} else {
		this.levels = append(this.levels, "N")
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.levels = []string{}
	this.helper(root)
	return fmt.Sprintf("%s", strings.Join(this.levels, ","))
}

func (this *Codec) connect() *TreeNode {
	if this.nodes[this.pos] == nil {
		this.pos += 1
		return nil
	}

	root := this.nodes[this.pos]
	this.pos += 1
	root.Left = this.connect()
	root.Right = this.connect()
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ss := strings.Split(data, ",")
	this.nodes = make([]*TreeNode, len(ss))

	for i, txt := range ss {
		switch txt {
		case "N":
			this.nodes[i] = nil
		default:
			v, _ := strconv.ParseInt(txt, 10, 32)
			this.nodes[i] = &TreeNode{
				Val: int(v),
			}
		}
		i++
	}

	this.pos = 0
	return this.connect()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
