package main

import (
	"testing"
)

func contains(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// testing mencari elemen yang ada di slice
func TestContains(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	if !contains(slice, 3) {
		t.Errorf("Expected slice to contain 3")
	}
	if contains(slice, 6) {
		t.Errorf("Expected slice not to contain 6")
	}
}
