package permuteUniquePack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	assert.Equal(t, 3, len(PermuteUnique([]int{1,1,2})))
}