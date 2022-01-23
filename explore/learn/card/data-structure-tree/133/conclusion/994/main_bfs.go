package explore

import (
	"errors"
	"fmt"
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
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("cannot pop from empty queue")
	}
	ret := q.buffer[q.head]
	q.head = (q.head + 1) % cap(q.buffer)
	q.count -= 1
	return ret, nil
}

func (q *Queue) Push(elem interface{}) error {
	if q.count == cap(q.buffer) {
		return errors.New("queue is full")
	}
	q.count += 1
	q.buffer[q.tail] = elem
	q.tail = (q.tail + 1) % cap(q.buffer)
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.buffer[q.head], nil
}

func (q *Queue) Size() int {
	return q.count
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
func connect_bfs(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := makeNewQueue(1000)
	_ = queue.Push(root)
	for !queue.IsEmpty() {
		size := queue.Size()
		fmt.Printf("size: %d\n", size)
		for i := 0; i < size; i++ {
			node, _ := queue.Pop()
			if i < size-1 {
				head, _ := queue.Peek()
				(node.(*Node)).Next = head.(*Node)
			}

			if (node.(*Node)).Left != nil {
				_ = queue.Push((node.(*Node)).Left)
			}
			if (node.(*Node)).Right != nil {
				_ = queue.Push((node.(*Node)).Right)
			}
		}
		fmt.Printf("size at end of loop: %d\n", queue.Size())
	}
	return root
}

func bfs(queue *Queue) {
	if queue.IsEmpty() {
		return
	}

	var nodes []*Node
	for !queue.IsEmpty() {
		node, _ := queue.Pop()
		nodes = append(nodes, node.(*Node))
	}

	prev := &Node{}
	for _, node := range nodes {
		prev.Next = node
		if node.Left != nil {
			_ = queue.Push(node.Left)
		}
		if node.Right != nil {
			_ = queue.Push(node.Right)
		}
		prev = node
	}
	bfs(queue)
}
