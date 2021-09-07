package romanToInt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	assert.Equal(t, 3, RomanToInt("III"))
}