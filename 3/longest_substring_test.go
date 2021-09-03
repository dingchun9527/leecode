package longestSubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	longest := LengthOfLongestSubstring("abcdabc")
	assert.Equal(t, 4, longest)
}
