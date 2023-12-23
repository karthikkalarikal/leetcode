package main

import "testing"

func TestNextGreatestLetter(t *testing.T) {
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
			name: "test1",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			want: 'c',
		},
		{
			name: "test2",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			want: 'f',
		},
		{
			name: "test3",
			args: args{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z'},
			want: 'x',
		},
		{
			name: "test4",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'},
			want: 'f',
		},
		{
			name: "test5",
			args: args{letters: []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, target: 'e'},
			want: 'n',
		},
		{
			name: "test6",
			args: args{letters: []byte{'e', 'e', 'e', 'k', 'q', 'q', 'q', 'v', 'v', 'v', 'y'}, target: 'e'},
			want: 'k',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("next greatest letter %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
