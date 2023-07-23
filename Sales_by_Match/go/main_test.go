package main

import (
	"testing"
)

// test case
func TestSockMerchant(t *testing.T) {
	n := int32(9)
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	expected := int32(3)

	actual := sockMerchant(n, ar)
	if actual != expected {
		t.Errorf("sockMerchant(%d, %v) = %d; expected %d", n, ar, actual, expected)
	}
}
