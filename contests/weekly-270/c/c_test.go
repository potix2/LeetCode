package main

import "testing"

const NULL = 0

func Test_getDirections(t *testing.T) {
	type args struct {
		root       *TreeNode
		startValue int
		destValue  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				root:       makeBinaryTree([]int{5, 1, 2, 3, NULL, 6, 4}),
				startValue: 3,
				destValue:  6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirections(tt.args.root, tt.args.startValue, tt.args.destValue); got != tt.want {
				t.Errorf("getDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}
