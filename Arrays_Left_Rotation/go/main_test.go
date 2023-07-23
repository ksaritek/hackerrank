package main

import (
	"reflect"
	"testing"
)

func Test_rotLeft(t *testing.T) {
	type args struct {
		a []int32
		d int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test Case 0",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				d: 4,
			},
			want: []int32{5, 1, 2, 3, 4},
		},
		{
			name: "Test Case 1",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				d: 2,
			},
			want: []int32{3, 4, 5, 1, 2},
		},
		{
			name: "Test Case 2",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				d: 5,
			},
			want: []int32{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotLeft(tt.args.a, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
