package testing

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddMultipleCases(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestAddEdgeCases(t *testing.T) {
	if Add(0, 0) != 0 {
		t.Error("Expected Add(0, 0) to be 0")
	}
	if Add(-1, 1) != 0 {
		t.Error("Expected Add(-1, 1) to be 0")
	}
}

func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3}
	if !contains(slice, 2) {
		t.Errorf("Expected slice to contain 2")
	}
	if contains(slice, 4) {
		t.Errorf("Expected slice not to contain 4")
	}
}

func WriteToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
					return "", err
	}
	return string(data), nil
}

func TestFileIO(t *testing.T) {
	filename := "testfile.txt"
	content := "Hello, Go!"

	// Menulis ke file
	err := WriteToFile(filename, content)
	if err != nil {
					t.Fatalf("Expected no error, got %v", err)
	}

	// Membaca dari file
	readContent, err := ReadFromFile(filename)
	if err != nil {
					t.Fatalf("Expected no error, got %v", err)
	}

	if readContent != content {
					t.Errorf("Expected %s, got %s", content, readContent)
	}

	// Bersihkan file setelah tes
	os.Remove(filename)
}