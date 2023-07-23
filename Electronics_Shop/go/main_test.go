package main

import "testing"

func Test_getMoneySpent(t *testing.T) {
	type args struct {
		keyboards []int32
		drives    []int32
		b         int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 0",
			args: args{
				keyboards: []int32{3, 1},
				drives:    []int32{5, 2, 8},
				b:         10,
			},
			want: 9,
		},
		{
			name: "Test Case 1",
			args: args{
				keyboards: []int32{4},
				drives:    []int32{5},
				b:         5,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneySpent(tt.args.keyboards, tt.args.drives, tt.args.b); got != tt.want {
				t.Errorf("getMoneySpent() = %v, want %v", got, tt.want)
			}
		})
	}
}
