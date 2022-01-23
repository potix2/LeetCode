package explore

import (
	"reflect"
	"testing"
)

func Test_connect_bfs(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				root: &Node{
					Val: 1,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val: 4,
						},
						Right: &Node{
							Val: 5,
						},
					},
					Right: &Node{
						Val: 3,
						Left: &Node{
							Val: 6,
						},
						Right: &Node{
							Val: 7,
						},
					},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Example2",
			args: args{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := connect_bfs(tt.args.root)
			traverse := []int{got.Val}
			for l := got.Left; l != nil; l = l.Left {
				for n := l; n != nil; n = n.Next {
					traverse = append(traverse, n.Val)
				}
			}
			if !reflect.DeepEqual(traverse, tt.want) {
				t.Errorf("connect() = %v, want %v", traverse, tt.want)
			}
		})
	}
}
