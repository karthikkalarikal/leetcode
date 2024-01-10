package main

import (
	"reflect"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}

	tests := []struct {
		name  string
		args  args
		want1 []int
		want2 []int
	}{
		{
			"test1",
			args{[]int{1, 1}, []int{2, 2}},
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"test2",
			args{[]int{1, 2}, []int{2, 3}},
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"test3",
			args{[]int{2}, []int{1, 3}},
			[]int{2, 3},
			[]int{3, 2},
		},
		{
			"test4",
			args{[]int{1, 2, 5}, []int{2, 4}},
			[]int{5, 4},
			[]int{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FairCandySwap(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want1) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("peak index  %v, want1 %v , want2 %v", got, tt.want1, tt.want2)
			}
		})
	}

}
