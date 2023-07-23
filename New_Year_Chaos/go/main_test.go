package main

import (
	"math"
	"testing"
)

func Test_calculateBribes(t *testing.T) {
	type args struct {
		q []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 1",
			args: args{
				q: []int32{2, 1, 5, 3, 4},
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				q: []int32{2, 5, 1, 3, 4},
			},
			want: math.MinInt32,
		},
		{
			name: "Test Case 3",
			args: args{
				q: []int32{1, 2, 5, 3, 7, 8, 6, 4},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateBribes(tt.args.q); got != tt.want {
				t.Errorf("calculateBribes() = %v, want %v", got, tt.want)
			}
		})
	}
}
