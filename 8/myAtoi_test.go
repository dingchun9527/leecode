package myAtoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 123, MyAtoi("123"))
}
