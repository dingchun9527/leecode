package nextPermutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1,2,3}
	NextPermutation(nums)
	assert.Equal(t, 1, nums[0])
	assert.Equal(t, 3, nums[1])
	assert.Equal(t, 2, nums[2])
}