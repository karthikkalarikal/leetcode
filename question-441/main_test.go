package main

import "testing"

func TestArrangeCoins(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{num: 5},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{num: 8},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{num: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrangeCoins(tt.args.num); got != tt.want {
				t.Errorf("ArrageCoins() %v, want %v", got, tt.want)
			}
		})
	}
}
