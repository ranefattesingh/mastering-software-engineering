package arrays_test

import (
	"testing"

	"github.com/ranefattesingh/interview-preparation/golang/final450/arrays"
	"github.com/stretchr/testify/assert"
)

func TestFindKthSmallestElement(t *testing.T) {
	testCases := []struct {
		inputArray     []int
		inputK         int
		expectedOutput int
	}{
		// {
		// 	inputArray:     []int{10, 5, 4, 3, 48, 6, 2, 33, 53, 10},
		// 	inputK:         4,
		// 	expectedOutput: 5,
		// },
		{
			inputArray:     []int{7, 10, 4, 3, 20, 15},
			inputK:         3,
			expectedOutput: 7,
		},
	}

	for _, testCase := range testCases {
		actualOutput := arrays.FindKthSmallestElement(testCase.inputArray, testCase.inputK)
		assert.Equal(t, testCase.expectedOutput, actualOutput)
	}
}
