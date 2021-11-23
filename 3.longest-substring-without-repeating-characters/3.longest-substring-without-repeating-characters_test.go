package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Example4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Space",
			args: args{
				s: " ",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
			t.Errorf("%q. lengthOfLongestSubstring() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
