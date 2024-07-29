package main

import "testing"

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			"empty string",
			"",
			map[string]int{},
		},
		{
			"single word",
			"hello",
			map[string]int{"hello": 1},
		},
		{
			"multiple words",
			"the >?quick brown) fox jumps*() over the lazy dog#^",
			map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := WordFrequency(tt.input)
			if !compareMaps(actual, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func compareMaps(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true

}
func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single letter",
			input:    "a",
			expected: true,
		},
		{
			name:     "single word1",
			input:    "racecar.",
			expected: true,
		},
		{
			name:     "single word2",
			input:    "racecars",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsPalindrome(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
