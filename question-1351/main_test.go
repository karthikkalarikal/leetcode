package main

import "testing"

func TestCountNeagetives(t *testing.T) {
	type args struct {
		grid [][]int
	}

	test := []struct {
		name string
		args args
		want int
	}{
		{
			"Test1",
			args{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}},
			8,
		},
		{
			"Test2",
			args{[][]int{{3, 2}, {1, 0}}},
			0,
		},
		{
			"Test3",
			args{[][]int{{5, 1, 0}, {-5, -5, -5}}},
			3,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNegatives(tt.args.grid); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}

}
