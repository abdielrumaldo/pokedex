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
			input:    "  hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HELLO   World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  do YoU    belive   In magic  ",
			expected: []string{"do", "you", "belive", "in", "magic"},
		},
	}

	for _, c := range cases {

		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length Test Length Check:\nInput: %v\nActual: %v\nExpected:%v\n", c.input, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Content Test Failed at Index %v:\nInput: %v\nActual: '%v'\nExpected '%v'\n", i, c.input, word, expectedWord)
			}

		}
	}
}
