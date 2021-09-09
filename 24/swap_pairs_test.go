package swapPairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	assert.Equal(t, 1, SwapPairs(&ListNode{Val: 1}).Val)
}
