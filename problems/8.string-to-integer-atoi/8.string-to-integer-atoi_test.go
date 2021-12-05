package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				"42",
			},
			want: 42,
		},
		{
			name: "Example2",
			args: args{
				"    -42",
			},
			want: -42,
		},
		{
			name: "Example3",
			args: args{
				"4193 with words",
			},
			want: 4193,
		},
		{
			name: "Example4",
			args: args{
				"words adn 987",
			},
			want: 0,
		},
		{
			name: "Example5",
			args: args{
				"-91283472332",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		if got := myAtoi(tt.args.s); got != tt.want {
			t.Errorf("%q. myAtoi() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
