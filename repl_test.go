package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "          test Case",
			expected: []string{"test", "case"},
		},
		{
			input:    "    test case AGAIN        ",
			expected: []string{"test", "case", "again"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected number of words: %v, \ngot %v", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word: %v, \ngot %v", c.expected[i], actual[i])
			}
		}
	}
}
