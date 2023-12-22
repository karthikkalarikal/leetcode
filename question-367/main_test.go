package main

import (
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{num: 2},
			want: false,
		},
		{
			name: "Case 2",
			args: args{num: 1},
			want: true,
		},
		{
			name: "Case 3",
			args: args{num: 4},
			want: true,
		},
		{
			name: "Case 5",
			args: args{num: 16},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() %v, want %v", got, tt.want)
			}
		})
	}
}
