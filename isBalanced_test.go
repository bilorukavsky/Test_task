package main

import (
	"fmt"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"{[()]}", true},
		{"{[(])}", false},
		{"{{[[(())]]}}", true},
		{"[()]", true},
		{"()", true},
		{"[}", false},
		{"(", false},
		{"[()]", true},
		{"", true},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx+1), func(t *testing.T) {
			result := isBalanced(tc.input)
			if result != tc.expected {
				t.Errorf("For input %s, expected %t but got %t", tc.input, tc.expected, result)
			}
		})
	}
}
