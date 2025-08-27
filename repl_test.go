package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "TODAY Is A gooD DAY",
			expected: []string{"today", "is", "a", "good", "day"},
		},
		{
			input:    "",
			expected: []string{},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
			continue
		}
		// Compare each element in the slices
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s', expected word '%s' at index %d but got '%s'", c.input, expectedWord, i, word)
			}
		}
	}
}
