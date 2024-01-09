package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[]int{1, 2, 2, 1}, []int{2, 2}},
			[]int{2},
		},
		{
			"test2",
			args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			[]int{4, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.num1, tt.args.num2); reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInsert() %v, want %v", got, tt.want)
			}
		})
	}
}
