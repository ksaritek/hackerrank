package main

import "testing"

func Test_sherlockAndAnagrams(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 0",
			args: args{
				s: "abba",
			},
			want: 4,
		},
		{
			name: "Test Case 1",
			args: args{
				s: "abcd",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sherlockAndAnagrams(tt.args.s); got != tt.want {
				t.Errorf("sherlockAndAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
