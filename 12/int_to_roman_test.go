package intToRoman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "IV", IntToRoman(4))
}