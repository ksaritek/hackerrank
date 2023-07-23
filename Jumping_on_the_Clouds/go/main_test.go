package main

import "testing"

func Test_jumpingOnClouds(t *testing.T) {
	type args struct {
		c []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 0",
			args: args{
				c: []int32{0, 0, 1, 0, 0, 1, 0},
			},
			want: 4,
		},
		{
			name: "Test Case 1",
			args: args{
				c: []int32{0, 0, 0, 0, 1, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpingOnClouds(tt.args.c); got != tt.want {
				t.Errorf("jumpingOnClouds() = %v, want %v", got, tt.want)
			}
		})
	}
}
