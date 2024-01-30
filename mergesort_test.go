package mergesort_test

import (
	"reflect"
	"sort"
	"testing"

	"modules.cedoromal.com/mergesort"
)

func TestMergeSort(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mergesort.Sort(tt.input)
			if !sort.IntsAreSorted(actual) {
				t.Errorf("Expected sorted slice, got %v", actual)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
