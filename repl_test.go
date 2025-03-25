package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic two words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "extra spaces",
			input:    "hello    world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "leading and trailing spaces",
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "hello",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", c.name, c.expected, actual)
			}
		})
	}
}
