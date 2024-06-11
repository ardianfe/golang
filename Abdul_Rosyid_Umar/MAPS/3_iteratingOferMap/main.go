package main

import "fmt"

func main() {
	m := map[string]any{
		"umurandi":  45,
		"umursinta": 17.5,
	}
	for key, val := range m {
		fmt.Printf("key: %s, val: %v\n", key, val)
	}
}
