package main

import (
	"testing"
)

func TestCheckCansplitArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Can split array",
			input:    []int{1, 3, 3, 4, 3},
			expected: 1,
		},
		{
			name:     "Cannot split array",
			input:    []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single element array",
			input:    []int{10},
			expected: 0,
		},
		{
			name:     "Array with equal halves",
			input:    []int{2, 2, 4, 2, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkCansplitArray(tt.input)
			if result != tt.expected {
				t.Errorf("checkCansplitArray(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
