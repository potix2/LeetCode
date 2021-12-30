package main

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example3",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: [][]int{[]int{1}},
		},
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: [][]int{{3}, {9,20}, {15,7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type fields struct {
		buffer []interface{}
		head   int
		tail   int
		count  int
	}
	type args struct {
		elem int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Push an elem to empty queue",
			fields: fields{
				buffer: []interface{}{0, 0, 0, 0, 0},
				head: 0,
				tail: 0,
				count: 0,
			},
			args: args{
				1,
			},
			wantErr: false,
		},
		{
			name: "Push an elem to the queue which has an element",
			fields: fields{
				buffer: []interface{}{1, 0, 0, 0, 0},
				head: 0,
				tail: 1,
				count: 1,
			},
			args: args{
				2,
			},
			wantErr: false,
		},
		{
			name: "Push an elem to the queue which has elements",
			fields: fields{
				buffer: []interface{}{1, 2, 3, 4, 0},
				head: 0,
				tail: 4,
				count: 4,
			},
			args: args{
				5,
			},
			wantErr: false,
		},
		{
			name: "Failed to push when the queue is full",
			fields: fields{
				buffer: []interface{}{1, 2, 3, 4, 5},
				head: 0,
				tail: 0,
				count: 5,
			},
			args: args{
				6,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				buffer: tt.fields.buffer,
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				count:  tt.fields.count,
			}
			if err := q.Push(tt.args.elem); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	type fields struct {
		buffer []interface{}
		head   int
		tail   int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "Pop from fully filled queue",
			fields: fields{
				buffer: []interface{}{1,2,3,4,5},
				head: 0,
				tail: 0,
				count: 5,
			},
			want: 1,
			wantErr: false,
		},
		{
			name: "Pop when the head point to the end of buffer",
			fields: fields{
				buffer: []interface{}{1,2,3,4,5},
				head: 4,
				tail: 4,
				count: 5,
			},
			want: 5,
			wantErr: false,
		},
		{
			name: "Failed to pop when the queue is empty",
			fields: fields{
				buffer: []interface{}{0, 0, 0, 0, 0},
				head: 0,
				tail: 0,
				count: 0,
			},
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				buffer: tt.fields.buffer,
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				count:  tt.fields.count,
			}
			got, err := q.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}