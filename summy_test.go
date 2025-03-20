package summy

import "testing"

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3, 4, 5})
	expected := 15

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
