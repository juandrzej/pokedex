package main

import "testing"

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
			input:    "your      momma",
			expected: []string{"your", "momma"},
		},

		{
			input:    "i	like_trains",
			expected: []string{"i", "like_trains"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expectedLen := len(c.expected)
		if len(actual) != expectedLen {
			t.Errorf("Expected and actual lenghts differ.")
			t.Fail()
		}

		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected and actual words differ.")
				t.Fail()
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
