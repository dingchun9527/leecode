package threeSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	results := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	assert.Equal(t, 2, len(results))
}
