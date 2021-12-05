package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "Example2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "Example3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Example4",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "Match all",
			args: args{
				s: "abba",
			},
			want: "abba",
		},
		{
			name: "repeat pattern",
			args: args{
				s: "xababa",
			},
			want: "ababa",
		},
	}
	for _, tt := range tests {
		if got := longestPalindrome(tt.args.s); got != tt.want {
			t.Errorf("%q. longestPalindrome() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
