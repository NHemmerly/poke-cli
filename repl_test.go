package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "		hello	 world		",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "	",
			expected: []string{},
		},
		{
			input:    "heLLo world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "		hello	 world		of   Pokemon",
			expected: []string{"hello", "world", "of", "pokemon"},
		},
		{
			input:    "フシギダネ ヒトカゲ",
			expected: []string{"フシギダネ", "ヒトカゲ"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Cases do not match")
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word and expected word are not the same %s", c)
				t.FailNow()
			}

		}
	}
}
