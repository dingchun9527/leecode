package isValidBrackets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, IsValid("()"))
}
