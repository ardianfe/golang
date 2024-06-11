package main

import (
	"testing"
)

func TestInventory(t *testing.T) {
	inventory := map[string]int{
		"pulpen": 1,
		"buku":   2,
	}

	if inventory["pulpen"] != 1 {
		t.Errorf("expected 1, got %v", inventory["pulpen"])
	}
	if inventory["buku"] != 2 {
		t.Errorf("expected 2, got %v", inventory["buku"])
	}
	if _, exists := inventory["penghapus"]; exists {
		t.Errorf("expected false, got %v", exists)
	}
}
