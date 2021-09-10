package strStr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, 1, StrStr("abcdef", "bcdef"), 1)
}
