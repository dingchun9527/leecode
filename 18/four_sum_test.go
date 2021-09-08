package fourSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFourSum(t *testing.T) {
	assert.Equal(t, 3, len(FourSum([]int{1, 0, -1, 0, -2, 2}, 0)))
}
