package main

type ListNode struct {
	Val int
	Next *ListNode
}

func makeNodes(list []int) *ListNode {
	var head *ListNode
	current := &head

	for _, a := range list {
		*current = &ListNode{
			Val: a,
			Next: nil,
		}
		current = &(*current).Next
	}
	return head
}

func deleteMiddle(head *ListNode) *ListNode {
	ints := make([]int, 100000)
	i := 0
	for cur := &head; *cur != nil; i++ {
		ints[i] = (*cur).Val
		cur = &((*cur).Next)
	}
	if i == 1 {
		return nil
	}

	midIndex := i / 2.0
	return makeNodes(append(ints[:midIndex], ints[midIndex+1:i]...))
}