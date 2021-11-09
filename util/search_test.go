package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchFirstGreater(t *testing.T) {
	arr := []int{2, 3, 10, 22}
	assert.Equal(t, 2, BinarySearchFirstGreater(arr, 5))
	assert.Equal(t, 2, BinarySearchFirstGreater(arr, 3))
	assert.Equal(t, 3, BinarySearchFirstGreater(arr, 22))
	assert.Equal(t, -1, BinarySearchFirstGreater(arr, 23))
}
