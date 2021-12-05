package main

import (
	"reflect"
	"testing"
)

func Test_findEvenNumbers(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{[]int{2,1,3,0}},
			want: []int{102,120,130,132,210,230,302,310,312,320},
		},
		{
			name: "Example2",
			args: args{[]int{2,2,8,8,2}},
			want: []int{ 222,228,282,288,822,828,882},
		},
		{
			name: "Example3",
			args: args{[]int{3,7,5}},
			want: []int{},
		},
		{
			name: "Example4",
			args: args{[]int{0,2,0,0}},
			want: []int{ 200},
		},
		{
			name: "Example5",
			args: args{[]int{0,0,0}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEvenNumbers(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
