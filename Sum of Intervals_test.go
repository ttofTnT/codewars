package codewars

import (
	"testing"
)

func TestSumOfIntervals(t *testing.T) {
	type args struct {
		intervals [][2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				intervals: [][2]int{{1, 2}, {6, 10}, {11, 15}},
			},
			want: 9,
		},
		{
			name: "test2",
			args: args{
				intervals: [][2]int{{1, 5}, {10, 20}, {1, 5}, {20, 30}},
			},
			want: 24,
		},
		{
			name: "test3",
			args: args{
				intervals: [][2]int{{1, 6}, {3, 20}, {20, 30}},
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("SumOfIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
