package main

import (
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args{makeNodes([]int{1,3,4,7,1,2,6})},
			want: makeNodes([]int{1,3,4,1,2,6}),
		},
		{
			name: "Example2",
			args: args{makeNodes([]int{1,2,3,4})},
			want: makeNodes([]int{1,2,4}),
		},
		{
			name: "Example3",
			args: args{makeNodes([]int{2,1})},
			want: makeNodes([]int{2}),
		},
		{
			name: "Example4",
			args: args{makeNodes([]int{2})},
			want: makeNodes([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
