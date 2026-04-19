package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 4 {
		t.Errorf("Add(2, 2) = %d; want 4", result)
	}
}
