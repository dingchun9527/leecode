package letterCombinations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	assert.Equal(t, 9, len(LetterCombinations("23")))
}
