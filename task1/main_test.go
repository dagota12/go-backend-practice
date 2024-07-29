package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	testCases := []struct {
		name     string
		subjects map[string]float32
		expected float32
	}{
		{
			name:     "empty subjects",
			subjects: map[string]float32{},
			expected: 0,
		},
		{
			name:     "single subject",
			subjects: map[string]float32{"Math": 90},
			expected: 90,
		},
		{
			name:     "multiple subjects",
			subjects: map[string]float32{"Applied": 90, "DSA": 85, "English": 78},
			expected: 84.33,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := calculateAverage(tc.subjects)
			const delta float64 = 0.01
			diff := math.Abs(float64(tc.expected - actual))
			fmt.Println(diff, delta)
			if diff > delta {
				t.Errorf("expected %.2f, got %.2f", tc.expected, actual)
			}
		})
	}
}
