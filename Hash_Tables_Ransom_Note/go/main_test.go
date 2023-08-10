package main

import "testing"

func Test_checkMagazine(t *testing.T) {
	type args struct {
		magazine []string
		note     []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 0",
			args: args{
				magazine: []string{"give", "me", "one", "grand", "today", "night"},
				note:     []string{"give", "one", "grand", "today"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkMagazine(tt.args.magazine, tt.args.note)
		})
	}
}
