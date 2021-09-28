package jump

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJump(t *testing.T) {
	assert.Equal(t, 2, Jump([]int{ 2,3,1,1,4}))
}