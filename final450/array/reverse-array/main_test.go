package main_test

import (
	"reflect"
	"testing"

	main "github.com/ranefattesingh/mastering-software-engineering/final450/array/reverse-array"
)

func TestReverseArray(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		size      int
		expect    []int
		expectErr error
	}{
		{
			name:      "Example1",
			input:     []int{1, 4, 3, 2, 6, 5},
			size:      6,
			expect:    []int{5, 6, 2, 3, 4, 1},
			expectErr: nil,
		},
		{
			name:      "Example2",
			input:     []int{4, 5, 1, 2},
			size:      4,
			expect:    []int{2, 1, 5, 4},
			expectErr: nil,
		},
		{
			name:      "Empty array",
			input:     []int{},
			size:      0,
			expect:    []int{},
			expectErr: main.ErrInvalidArraySize,
		},
		{
			name:      "Negative size",
			input:     []int{1, 2, 3},
			size:      -1,
			expect:    []int{1, 2, 3},
			expectErr: main.ErrInvalidArraySize,
		},
		{
			name:      "Single element",
			input:     []int{42},
			size:      1,
			expect:    []int{42},
			expectErr: nil,
		},
		{
			name:      "Odd number of elements",
			input:     []int{1, 2, 3, 4, 5},
			size:      5,
			expect:    []int{5, 4, 3, 2, 1},
			expectErr: nil,
		},
		{
			name:      "Size less than array length",
			input:     []int{1, 2, 3, 4, 5},
			size:      3,
			expect:    []int{3, 2, 1, 4, 5},
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrCopy := make([]int, len(tt.input))
			copy(arrCopy, tt.input)
			err := main.ReverseArray(arrCopy, tt.size)
			if err != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
			if !reflect.DeepEqual(arrCopy, tt.expect) {
				t.Errorf("expected array %v, got %v", tt.expect, arrCopy)
			}
		})
	}
}
