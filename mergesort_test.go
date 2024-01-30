package mergesort_test

import (
	"reflect"
	"testing"

	"modules.cedoromal.com/mergesort"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element slice",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "sorted slice",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "reversed slice",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "random slice",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:     "random slice with negative",
			input:    []int{3, -1, 4, 1, 5, 9, 2, 6, 5, -3, 5},
			expected: []int{-3, -1, 1, 2, 3, 4, 5, 5, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := mergesort.Sort(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestSortReversed(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element slice",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "sorted slice",
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "reversed slice",
			input:    []int{3, 2, 1},
			expected: []int{3, 2, 1},
		},
		{
			name:     "random slice",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{9, 6, 5, 5, 5, 4, 3, 3, 2, 1, 1},
		},
		{
			name:     "random slice with negative",
			input:    []int{3, -1, 4, 1, 5, 9, 2, 6, 5, -3, 5},
			expected: []int{9, 6, 5, 5, 5, 4, 3, 2, 1, -1, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := mergesort.SortReversed(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestSortFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "empty slice",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "single element slice",
			input:    []float64{0.1},
			expected: []float64{0.1},
		},
		{
			name:     "sorted slice",
			input:    []float64{0.1, 0.2, 0.3},
			expected: []float64{0.1, 0.2, 0.3},
		},
		{
			name:     "reversed slice",
			input:    []float64{0.3, 0.2, 0.1},
			expected: []float64{0.1, 0.2, 0.3},
		},
		{
			name:     "random slice",
			input:    []float64{0.3, 0.1, 0.4, 0.1, 0.5, 0.9, 0.2, 0.6, 0.5, 0.3, 0.5},
			expected: []float64{0.1, 0.1, 0.2, 0.3, 0.3, 0.4, 0.5, 0.5, 0.5, 0.6, 0.9},
		},
		{
			name:     "random slice with negative",
			input:    []float64{0.3, -0.1, 0.4, 0.1, 0.5, 0.9, 0.2, 0.6, 0.5, -0.3, 0.5},
			expected: []float64{-0.3, -0.1, 0.1, 0.2, 0.3, 0.4, 0.5, 0.5, 0.5, 0.6, 0.9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := mergesort.Sort(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestSortFloatReversed(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "empty slice",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "single element slice",
			input:    []float64{0.1},
			expected: []float64{0.1},
		},
		{
			name:     "sorted slice",
			input:    []float64{0.1, 0.2, 0.3},
			expected: []float64{0.3, 0.2, 0.1},
		},
		{
			name:     "reversed slice",
			input:    []float64{0.3, 0.2, 0.1},
			expected: []float64{0.3, 0.2, 0.1},
		},
		{
			name:     "random slice",
			input:    []float64{0.3, 0.1, 0.4, 0.1, 0.5, 0.9, 0.2, 0.6, 0.5, 0.3, 0.5},
			expected: []float64{0.9, 0.6, 0.5, 0.5, 0.5, 0.4, 0.3, 0.3, 0.2, 0.1, 0.1},
		},
		{
			name:     "random slice with negative",
			input:    []float64{0.3, -0.1, 0.4, 0.1, 0.5, 0.9, 0.2, 0.6, 0.5, -0.3, 0.5},
			expected: []float64{0.9, 0.6, 0.5, 0.5, 0.5, 0.4, 0.3, 0.2, 0.1, -0.1, -0.3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := mergesort.SortReversed(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
