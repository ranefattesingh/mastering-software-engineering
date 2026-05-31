package main_test

import (
	"errors"
	"testing"

	main "github.com/ranefattesingh/mastering-software-engineering/final450/array/minimum-and-maximum"
)

func TestFindMinimumAndMaximum(t *testing.T) {
	testCases := []struct {
		name    string
		arr     []int
		size    int
		wantMin int
		wantMax int
		wantErr error
	}{
		{
			name:    "Example1",
			arr:     []int{3, 5, 4, 1, 9},
			size:    5,
			wantMin: 1,
			wantMax: 9,
			wantErr: nil,
		},
		{
			name:    "Example2",
			arr:     []int{22, 14, 8, 17, 35, 3},
			size:    6,
			wantMin: 3,
			wantMax: 35,
			wantErr: nil,
		},
		{
			name:    "SingleElement",
			arr:     []int{42},
			size:    1,
			wantMin: 42,
			wantMax: 42,
			wantErr: nil,
		},
		{
			name:    "AllSame",
			arr:     []int{7, 7, 7, 7},
			size:    4,
			wantMin: 7,
			wantMax: 7,
			wantErr: nil,
		},
		{
			name:    "NegativeNumbers",
			arr:     []int{-5, -1, -9, -3},
			size:    4,
			wantMin: -9,
			wantMax: -1,
			wantErr: nil,
		},
		{
			name:    "MixedNumbers",
			arr:     []int{-10, 0, 10, 5, -5},
			size:    5,
			wantMin: -10,
			wantMax: 10,
			wantErr: nil,
		},
		{
			name:    "EmptyArray",
			arr:     []int{},
			size:    0,
			wantMin: -1,
			wantMax: -1,
			wantErr: main.ErrInvalidArraySize,
		},
		{
			name:    "ZeroSizeNonEmptyArray",
			arr:     []int{1, 2, 3},
			size:    0,
			wantMin: -1,
			wantMax: -1,
			wantErr: main.ErrInvalidArraySize,
		},
		{
			name:    "NegativeSize",
			arr:     []int{1, 2, 3},
			size:    -1,
			wantMin: -1,
			wantMax: -1,
			wantErr: main.ErrInvalidArraySize,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			min, max, err := main.FindMinimumAndMaximum(tc.arr, tc.size)
			if min != tc.wantMin || max != tc.wantMax {
				t.Errorf("%s: got min=%d, max=%d; want min=%d, max=%d", tc.name, min, max, tc.wantMin, tc.wantMax)
			}
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("%s: got err=%v; want err=%v", tc.name, err, tc.wantErr)
			}
		})
	}
}
