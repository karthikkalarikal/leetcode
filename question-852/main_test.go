package main

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{0, 1, 0}},
			1,
		},
		{
			"test2",
			args{[]int{0, 2, 1, 0}},
			1,
		},
		{
			"test3",
			args{[]int{0, 10, 5, 2}},
			1,
		},
		{
			"test4",
			args{[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PeakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peak index  %v, want %v", got, tt.want)
			}
		})
	}

}
