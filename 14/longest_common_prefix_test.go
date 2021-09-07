package longestCommonPrex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", LongestCommonPrefix([]string{"flower","flow","flight"}))
}