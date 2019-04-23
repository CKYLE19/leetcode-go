package leetcode

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Base Case",
			args: args{s: "abcabcbb"},
			want: "abc",
		},
		{
			name: "Only one char",
			args: args{s: "bbbbb"},
			want: "b",
		},
		{
			name: "In middle of string",
			args: args{s: "pwwkew"},
			want: "wke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
