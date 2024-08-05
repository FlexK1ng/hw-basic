package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "Simple case",
			input:    "hello world hello",
			expected: map[string]int{"hello": 2, "world": 1},
		},
		{
			name:     "Case sensitivity",
			input:    "Hello hello HeLLo",
			expected: map[string]int{"hello": 3},
		},
		{
			name:     "Punctuation",
			input:    "Hello, world! Hello...",
			expected: map[string]int{"hello": 2, "world": 1},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "Single word",
			input:    "word",
			expected: map[string]int{"word": 1},
		},
		{
			name:     "Multiple spaces",
			input:    "word   word",
			expected: map[string]int{"word": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countWords(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, actual)
			}
		})
	}
}
