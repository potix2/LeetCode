package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func makeListNode(list []int) *ListNode {
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

func toStr(l *ListNode) string {
	b := strings.Builder{}
	fmt.Fprint(&b, "[")
	for c := l; c != nil; c = c.Next {
		fmt.Fprintf(&b, " %d", c.Val)
	}
	fmt.Fprint(&b, " ]")
	return b.String()
}

func Test_makeListNode(t *testing.T) {
	actual := toStr(makeListNode([]int{}))
	expected := "[ ]"
	if actual != expected {
		t.Errorf("expect=%v, but got = %v\n", expected, actual)
	}

	actual = toStr(makeListNode([]int{3, 2, 1}))
	expected = "[ 3 2 1 ]"
	if actual != expected {
		t.Errorf("expect=%v, but got = %v\n", expected, actual)
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args {
				l1: makeListNode([]int{ 2, 4, 3}),
				l2: makeListNode([]int{ 5, 6, 4}),
			},
			want: makeListNode([]int{ 7, 0, 8}),
		},
		{
			name: "Example2",
			args: args{
				l1: makeListNode([]int{0}),
				l2: makeListNode([]int{0}),
			},
			want: makeListNode([]int{0}),
		},
		{
			name: "Example3",
			args: args{
				l1: makeListNode([]int {9,9,9,9,9,9,9}),
				l2: makeListNode([]int {9,9,9,9}),
			},
			want: makeListNode([]int{8,9,9,9,0,0,0,1}),
		},
		{
			name: "Example4",
			args: args{
				l1: makeListNode([]int {9,9,9,9}),
				l2: makeListNode([]int {9,9,9,9,9,9,9}),
			},
			want: makeListNode([]int{8,9,9,9,0,0,0,1}),
		},
	}
	for _, tt := range tests {
		if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. addTwoNumbers() = %v, want %v", tt.name, toStr(got), toStr(tt.want))
		}
	}
}
