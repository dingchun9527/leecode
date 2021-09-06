package longestPalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	longest := LongestPalindrome("abcdefg")
	assert.Equal(t, "a", longest)
	longest = LongestPalindrome("abcdefgfed")
	assert.Equal(t, "defgfed", longest)

}