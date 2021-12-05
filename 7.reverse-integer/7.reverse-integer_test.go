package main

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				123,
			},
			want: 321,
		},
		{
			name: "Example2",
			args: args{
				-123,
			},
			want: -321,
		},
		{
			name: "Example3",
			args: args{
				120,
			},
			want: 21,
		},
		{
			name: "Example4",
			args: args{
				0,
			},
			want: 0,
		},
		{
			name: "OutOfRange",
			args: args{
				1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := reverse(tt.args.x); got != tt.want {
			t.Errorf("%q. reverse() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
