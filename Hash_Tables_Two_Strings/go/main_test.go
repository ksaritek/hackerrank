package main

import "testing"

func Test_twoStrings(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test case 00", args{"hello", "world"}, "YES"},
		{"Test case 01", args{"hi", "world"}, "NO"},
		{"Test case 02", args{"wouldyoulikefries", "abcabcabcabcabcabc"}, "NO"},
		{"Test case 03", args{"hackerrankcommunity", "cdecdecdecde"}, "YES"},
		{"Test case 04", args{"jackandjill", "wentupthehill"}, "YES"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoStrings(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("twoStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
