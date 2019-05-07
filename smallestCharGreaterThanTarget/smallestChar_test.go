package leetcode

import "testing"

func Test_smallestChar(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Base Case",
			args: args{
				letters: []byte{[]byte("c")[0], []byte("f")[0], []byte("j")[0]},
				target:  []byte("a")[0],
			},
			want: []byte("c")[0],
		},
		{
			name: "Next Highest: target in letters",
			args: args{
				letters: []byte{[]byte("c")[0], []byte("f")[0], []byte("j")[0]},
				target:  []byte("c")[0],
			},
			want: []byte("f")[0],
		},
		{
			name: "Next Highest: target not in letters",
			args: args{
				letters: []byte{[]byte("c")[0], []byte("f")[0], []byte("j")[0]},
				target:  []byte("d")[0],
			},
			want: []byte("f")[0],
		},
		{
			name: "Wrap Around",
			args: args{
				letters: []byte{[]byte("c")[0], []byte("f")[0], []byte("j")[0]},
				target:  []byte("k")[0],
			},
			want: []byte("c")[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestChar(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("smallestChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
