package plusTwoNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	array := []int{2, 7, 11, 15}
	target := 9
	index := TwoSum(array, target)
	assert.Equal(t, 2, len(index))
	assert.Equal(t, 0, index[0])
	assert.Equal(t, 1, index[1])
}
