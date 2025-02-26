package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "bingoBongoworlds",
			expected: []string{"bingobongoworlds"},
		},
		{
			input: " ",
			expected: []string{},
		},
		{
			input: "BingO BongO, I am in the CONGO",
			expected: []string{"bingo", "bongo,", "i", "am", "in", "the", "congo"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("function produced too few segments. Expected: %s . Actual: %s .", c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("function produced word which does not match expected word. Expected: %s . Actual: %s .", c.expected, actual)
			}
		}
	}
}
