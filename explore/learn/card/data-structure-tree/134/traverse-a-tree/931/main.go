package main

import (
	"errors"
)

type Queue struct {
	buffer []interface{}
	head   int
	tail   int
	count  int
}

func makeNewQueue(capacity int) *Queue {
	return &Queue{
		buffer: make([]interface{}, capacity),
		head: 0,
		tail: 0,
		count: 0,
	}
}

func (q *Queue)Pop() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("cannot pop from empty queue")
	}
	ret := q.buffer[q.head]
	q.head = (q.head + 1)  % cap(q.buffer)
	q.count -= 1
	return ret, nil
}

func (q *Queue)Push(elem interface{}) error {
	if q.count == cap(q.buffer) {
		return errors.New("queue is full")
	}
	q.count += 1
	q.buffer[q.tail] = elem
	q.tail = (q.tail + 1) % cap(q.buffer)
	return nil
}

func (q *Queue)IsEmpty() bool {
	return q.count == 0
}

func (q *Queue)Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.buffer[q.head], nil
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func bfs(queue *Queue) [][]int {
	if queue.IsEmpty() {
		return [][]int{}
	}

	var nodes []*TreeNode
	var elems []int
	for !queue.IsEmpty() {
		node, _ := queue.Pop()
		nodes = append(nodes, node.(*TreeNode))
		elems = append(elems, node.(*TreeNode).Val)
	}

	for _, node := range nodes {
		if node.Left != nil {
			_ = queue.Push(node.Left)
		}
		if node.Right != nil {
			_ = queue.Push(node.Right)
		}
	}
	return append([][]int{ elems }, bfs(queue)...)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := makeNewQueue(2000)
	_ = queue.Push(root)
	return bfs(queue)
}