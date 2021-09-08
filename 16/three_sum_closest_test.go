package threeSumClosest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	assert.Equal(t, -101, ThreeSumClosest([]int{-100,-98,-2,-1}, -101))
}