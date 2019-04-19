package leetcode

import "testing"

func Test_singleElement(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happyPath",
			args: args{
				arr: []int{
					1,
					1,
					2,
					3,
					3,
					4,
					4,
					8,
					8,
				},
			},
			want: 2,
		},
		{
			name: "noSingles",
			args: args{arr: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleElement(tt.args.arr); got != tt.want {
				t.Errorf("singleElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
