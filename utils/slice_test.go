package utils

import "testing"

func TestIndexOfElementIsFound(t *testing.T) {
	target := "item"
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{target}, 0},
		{[]string{target, "not", "not"}, 0},
		{[]string{"not", "not", target}, 2},
		{[]string{target, "not", target}, 0},
	}

	for _, test := range tests {
		index := IndexOf(test.input, target)

		if index != test.expected {
			t.Errorf("Expected index %d got %d", test.expected, index)
		}
	}
}

func TestIndexOfElementIsNotFound(t *testing.T) {
	target := "item"
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{}, -1},
		{[]string{"not", "not"}, -1},
		{nil, -1},
	}

	for _, test := range tests {
		index := IndexOf(test.input, target)

		if index != test.expected {
			t.Errorf("Expected index %d got %d", test.expected, index)
		}
	}
}

func TestSliceContainsReturnsTrue(t *testing.T) {
	target := "item"
	tests := [][]string{
		{target},
		{target, "not", "not"},
		{"not", "not", target},
		{target, "not", target},
	}

	for _, test := range tests {
		found := Contains(test, target)

		if !found {
			t.Errorf("Expected item to be found")
		}
	}
}

func TestSliceContainsReturnsFalse(t *testing.T) {
	target := "item"
	tests := [][]string{
		{},
		{"not", "not"},
		nil,
	}

	for _, test := range tests {
		found := Contains(test, target)

		if found {
			t.Errorf("Expected item to not be found")
		}
	}
}
