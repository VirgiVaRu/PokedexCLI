package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input string
		expected []string
	}{
		"extra spaces": {input: "  hello  world  ", expected: []string{"hello", "world"},},
		"upper cases" : {input: "HeLLO wORlD", expected: []string{"hello", "world"},},
		"extra spaces and upper cases": {input: "  HelLo  WorlD   ", expected: []string{"hello", "world"},},
	}

	for name, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of the answers does not match. Expected: %v Actual: %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s: The words at index %d do not match. Expected: %v Actual: %v", name, i, c.expected[i], actual[i])
			}
		}
	}


}