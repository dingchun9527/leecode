package reverseInt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse(123), 321)
}
