package leetcode

import "testing"

func Test_longestPrefix(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Base Case",
			args: args{
				words: []string{
					"hello",
					"hell",
					"heap",
					"heat",
					"her",
				},
			},
			want: "he",
		},
		{
			name: "Similar Case",
			args: args{
				words: []string{
					"hear",
					"heal",
					"heap",
					"heat",
				},
			},
			want: "hea",
		},
		{
			name: "No common",
			args: args{
				words: []string{
					"hello",
					"blah",
					"foo",
					"nah",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPrefix(tt.args.words); got != tt.want {
				t.Errorf("longestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
