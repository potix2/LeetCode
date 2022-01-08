package explore

import (
	"testing"
)

func Test_countUnivalSubtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnivalSubtrees(tt.args.root); got != tt.want {
				t.Errorf("countUnivalSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
