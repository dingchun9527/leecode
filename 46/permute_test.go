package permute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	assert.Equal(t, 6, len(Permute([]int{1,2,3})))
}