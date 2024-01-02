package main

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{1, 3, 5, 6}, 5},
			2,
		},
		{
			"test2",
			args{[]int{1, 3, 5, 6}, 2},
			1,
		},
		{
			"test3",
			args{[]int{1, 3, 5, 6}, 7},
			4,
		},
		{
			"test4",
			args{[]int{1, 3}, 2},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() %v, want %v", got, tt.want)
			}
		})
	}
}
