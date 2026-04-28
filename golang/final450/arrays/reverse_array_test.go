package arrays_test

import (
	"testing"

	"github.com/ranefattesingh/interview-preparation/golang/final450/arrays"
	"github.com/stretchr/testify/assert"
)

func TestReverseArray(t *testing.T) {
	input := []int{1, 4, 3, 2, 6, 5}

	expected := []int{5, 6, 2, 3, 4, 1}

	actual := arrays.ReverseArray(input)

	assert.Equal(t, expected, actual)

}
