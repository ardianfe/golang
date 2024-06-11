package main

import "testing"

func TestDivide(t *testing.T) {
	test := []struct {
		a, b float64
		want float64
		err  bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true},
		{7, 2, 3.5, false},
	}

	for _, tt := range test {
		got, err := Divide(tt.a, tt.b)
		if (err != nil) != tt.err {
			t.Errorf("Divide(%f, %f) error = %v, Err %v", tt.a, tt.b, err, tt.err)
		}
		if got != tt.want {
			t.Errorf("Divide(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
		}
	}
}
