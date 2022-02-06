package explore

import (
	"fmt"
	"reflect"
	"testing"
)

func printTreeNode(root *TreeNode, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", root.Val)
	if root.Left != nil {
		printTreeNode(root.Left, depth+1)
	}
	if root.Right != nil {
		printTreeNode(root.Right, depth+1)
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example1",
			args: args{
				"1,2,N,N,3,4,N,N,5,N,N",
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "Example2",
			args: args{
				"",
			},
			want: nil,
		},
		{
			name: "Example3",
			args: args{
				"-1,N,N",
			},
			want: &TreeNode{
				Val: -1,
			},
		},
		{
			name: "Example4",
			args: args{
				//"4 0,N,2,N,N,N,4,N,N,N,N,N,N,N,5",
				"0,N,2,N,4,N,5,N,N",
			},
			want: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Log("got=\n")
				printTreeNode(got, 0)
				t.Log("want=\n")
				printTreeNode(tt.want, 0)
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: "1,2,N,N,3,4,N,N,5,N,N",
		},
		{
			name: "Example2",
			args: args{
				root: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
