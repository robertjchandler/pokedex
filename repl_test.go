package main

import (
	"fmt"
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
			input:    " HELLO WORLD ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hELLO wORLD  ",
			expected: []string{"hello", "world"},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			failCount++
			t.Errorf(`---------------------------------
Actual length does not equal expected length.
Actual:     %v
Expected:  	%v
Fail
`, len(actual), len(c.expected))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Actual:     %v
Expected:  	%v
Pass
`, len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				failCount++
				t.Errorf(`---------------------------------
Word does not match expected word.
Word:			"%v"
Expected Word:	"%v"
Fail
`, word, expectedWord)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Word:			"%v"
Expected Word:	"%v"
Pass
`, word, expectedWord)
			}
		}
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
