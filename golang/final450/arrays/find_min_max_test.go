package arrays_test

import (
	"testing"

	"github.com/ranefattesingh/interview-preparation/golang/final450/arrays"
	"github.com/stretchr/testify/assert"
)

func TestFindMinMax(t *testing.T) {
	input := []int{1, 4, 3, 5, 8, 6}
	expected := []int{1, 8}

	actual := arrays.FindMinMax(input)

	assert.Equal(t, expected, actual)
}
